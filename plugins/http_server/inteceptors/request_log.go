package inteceptors

import (
	"bytes"
	"fmt"
	"io"
	"match-system/plugins"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	body        *bytes.Buffer
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		body:           bytes.NewBuffer(nil),
	}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (r *responseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger *plugins.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error("error", err)
				}
			}()

			startTime := time.Now()
			logger.Info(fmt.Sprintf("Started => %s %s", r.Method, r.URL.Path))

			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Error(fmt.Println("Error reading request body:", err))
			}
			r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			endTime := time.Now()
			logger.Info(fmt.Sprintf("Request Payload: %s", string(requestBody)))
			logger.Info(fmt.Sprintf("Response Body: %s", wrapped.body.String()))
			logger.Info(fmt.Sprintf("Completed %s in %v", r.URL.Path, endTime.Sub(startTime)))
		}

		return http.HandlerFunc(fn)
	}
}
