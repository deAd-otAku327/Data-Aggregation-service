package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type contextKey int8

const RequestIDKey contextKey = iota

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapResponseWriterWithLogging(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func Logging(log *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			rid := uuid.New().String()
			w.Header().Set("X-Request-ID", rid)

			log.Info(
				"Request:",
				slog.String("request_id", rid),
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
			)

			startReq := time.Now()

			wrappedrw := wrapResponseWriterWithLogging(w)
			next.ServeHTTP(wrappedrw, r.WithContext(context.WithValue(r.Context(), RequestIDKey, rid)))

			responseTime := time.Since(startReq).Milliseconds()

			log.Info(
				"Response:",
				slog.String("request_id", rid),
				slog.Int("status_code", wrappedrw.statusCode),
				slog.String("resp_time", strconv.Itoa(int(responseTime))),
			)
		})
	}
}
