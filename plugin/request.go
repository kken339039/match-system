package plugin

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(p []byte) (int, error) {
	r.body.Write(p)
	return r.ResponseWriter.Write(p)
}

func RequestInterceptor(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		logger.Info(fmt.Sprintf("Started => %s %s", r.Method, r.URL.Path))

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(fmt.Println("Error reading request body:", err))
		}
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		recorder := &responseRecorder{w, bytes.NewBuffer(nil)}
		next.ServeHTTP(recorder, r)

		endTime := time.Now()
		logger.Info(fmt.Sprintf("Request Payload: %s", string(requestBody)))
		logger.Info(fmt.Sprintf("Response Body: %s", recorder.body.String()))
		logger.Info(fmt.Sprintf("Completed %s in %v", r.URL.Path, endTime.Sub(startTime)))
	})
}
