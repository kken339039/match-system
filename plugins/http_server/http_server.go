package http_server

import (
	"encoding/json"
	"fmt"
	"match-system/plugins"
	"match-system/plugins/http_server/inteceptors"
	"net"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	listener *net.Listener
	router   *mux.Router
}

func (m HttpServer) Serve() {
	env := plugins.SysEnv
	logger := plugins.SysLogger

	var port string = env.GetEnv("SERVICE_PORT")
	if len(port) == 0 {
		port = "3000"
	}

	if m.listener == nil {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			logger.WithFields(logrus.Fields{"port": port, "error": err}).Error("Unable to listen to port")
		}
		m.listener = &lis
	}

	var allowOrigins []string
	if len(env.GetEnv("WHITELIST_DOMAINS")) > 0 {
		allowOrigins = strings.Split(env.GetEnv("WHITELIST_DOMAINS"), ",")
	} else {
		allowOrigins = []string{"*"}
	}

	cors := cors.New(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	})

	lis := *m.listener
	logger.Info(fmt.Sprintf("Http server is up and running on %s", lis.Addr().String()))
	m.router.MethodNotAllowedHandler = notAllowedHandler(m.router)
	loggingMiddleware := inteceptors.LoggingMiddleware(logger)
	loggedRouter := loggingMiddleware(m.router)
	if err := http.Serve(lis, cors.Handler(loggedRouter)); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func NewHttpServer(router *mux.Router) *HttpServer {
	s := &HttpServer{
		router: router,
	}
	return s
}

func EmptyResoponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func Resoponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	if reflect.ValueOf(payload).IsNil() {
		payload = make(map[string]string)
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)

	if err != nil {
		InternalServerError(w, r, err)
		return
	}
}

func notAllowedHandler(x *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := x.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			var routeMatch mux.RouteMatch
			if route.Match(r, &routeMatch) || routeMatch.MatchErr == mux.ErrMethodMismatch {
				m, _ := route.GetMethods()
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(m, ", "))
			}
			return nil
		})

		if err != nil {
			InternalServerError(w, r, err)
			return
		}
		NotFound(w, r)
	})
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	err := errors.New("Unauthorized")
	_, err = w.Write(formatErrorResponse(err))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	errors := errors.New("Forbidden")
	statusCode, err := w.Write(formatErrorResponse(errors))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	statusCode, err := w.Write(formatErrorResponse(err))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	statusCode, err := w.Write(formatErrorResponse(err))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := errors.New("Not Found")
	statusCode, err := w.Write(formatErrorResponse(err))
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
}

func formatErrorResponse(err error) []byte {
	errorRes := make(map[string]string)
	errorRes["errorMessage"] = err.Error()
	jsonBytes, _ := json.Marshal(errorRes)

	return jsonBytes
}
