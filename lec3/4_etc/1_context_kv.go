package main

import (
	"context"
	"net/http"
)

func main() {
	a := NewAuthorizer()
	h := WithAuth(a, http.HandlerFunc(Handle))
	http.ListenAndServe("/", h)
}

const TokenContextKey = "MyAppToken"

func WithAuth(a Authorizer, next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r) // continue without token
			return
		}

		token, err := a.Authorize(auth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), TokenContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if token := r.Context().Value(TokenContextKey); token != nil {
		// User is logged in
	} else {
		// User is not logged in
	}
}
