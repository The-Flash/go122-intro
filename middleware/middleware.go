package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

// MiddlewareStack chains multiple middlewares
func MiddlewareStack(ms ...Middleware) Middleware {
	return Middleware(func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}
		return next
	})
}

// wrappedResponseWriter add statusCode to ResponseWriter
type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggerWrapped(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Println(time.Since(start), r.Method, r.URL.Path, wrapped.statusCode)
	})
}

func IsLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking is logged in...")
		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking if is admin...")
		next.ServeHTTP(w, r)
	})
}

func IsCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("checking if is customer...")
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(start), r.Method, r.URL.Path)
	})
}
