package internal

import (
	"fmt"
	"net/http"

	"match-system/plugin"

	"github.com/gorilla/mux"
)

type Mux struct {
	router *mux.Router
}

func NewMux() *Mux {
	router := mux.NewRouter()
	return &Mux{
		router: router,
	}
}

func (m *Mux) Serve() {
	env := plugin.SysEnv
	logger := plugin.SysLogger
	port := fmt.Sprintf(":%s", env.GetEnv("PORT"))

	http.Handle("/", plugin.RequestInterceptor(m.router, logger))

	logger.Info(fmt.Sprintf("== Server is running on%s", port))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		logger.Error(fmt.Sprintf("== Failed to run server, error: %s", err))
	}
}
