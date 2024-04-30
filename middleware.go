package main

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func UserOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			encode(w, r, http.StatusUnauthorized, "Not authorised")
			return
		}
		tokenString = tokenString[len("Bearer "):]
		err := VerifyToken(tokenString, getenv)
		if err != nil {
			encode(w, r, http.StatusUnauthorized, "Not authorised")
			return
		}
		next.ServeHTTP(w, r)
	})
}
