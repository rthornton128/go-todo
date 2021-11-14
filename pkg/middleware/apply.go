package middleware

import "net/http"

func Apply(middlewares ...http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, h := range middlewares {
			h.ServeHTTP(w, r)
		}
	}
}
