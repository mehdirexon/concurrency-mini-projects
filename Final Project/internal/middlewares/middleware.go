package middlewares

import "net/http"

// Use adds a middleware to the chain.
func Use(mw func(http.Handler) http.Handler) {
	middlewares = append(middlewares, mw)
}

// Apply applies the middleware chain on the handler(mux handler)
func Apply(h http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// With registers certain middlewares only on one handler
func With(h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.HandlerFunc {
	var handler http.Handler = h
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler.ServeHTTP
}
