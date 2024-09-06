package middleware

import (
	"log/slog"
	"net/http"
)

type ResponseWriterMetaHolder struct {
	http.ResponseWriter
	Status int
}

func (wMeta *ResponseWriterMetaHolder) WriteHeader(statusCode int) {
	wMeta.ResponseWriter.WriteHeader(statusCode)
	wMeta.Status = statusCode
}

func LogIncomingRequests(targetFunc http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received Request", slog.String("method", r.Method), slog.String("uri", r.RequestURI))

		wMetaInterceptor := &ResponseWriterMetaHolder{
			ResponseWriter: w,
		}

		targetFunc(wMetaInterceptor, r)
		slog.Info("Responding Request",
			slog.String("method", r.Method),
			slog.String("uri", r.RequestURI),
			slog.Int("status_code", wMetaInterceptor.Status),
			slog.String("origin", r.RemoteAddr),
		)
	})
}

type middlewareFunc func(http.HandlerFunc) http.Handler

func CreateMiddleware(baseMiddleware middlewareFunc) *Middleware {
	return &Middleware{baseMiddleware}
}

type Middleware struct {
	baseMiddleware middlewareFunc
}

func (mw *Middleware) Apply(targetFunc http.HandlerFunc) http.Handler {
	return mw.baseMiddleware(targetFunc)
}
