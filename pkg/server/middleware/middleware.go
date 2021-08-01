package middleware

import (
	"context"
	"net/http"
)

type KeyRequest struct{}

// Middleware to be used for validation and context mapping
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), KeyRequest{}, r)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
