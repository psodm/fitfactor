package main

import (
	"context"
	"net/http"
)

func addRoutes(ctx context.Context, mux *http.ServeMux, userStore *UserStore) {
	mux.Handle("GET /healthz", handleHealthzPlease())
	mux.Handle("POST /api/v1/auth/login", handleLogin(ctx, userStore))
	mux.Handle("GET /api/v1/users/{id}", handleFindUser(ctx, userStore))
	mux.Handle("POST /api/v1/users", handleCreateUserAccount(ctx, userStore))
	mux.Handle("GET /", http.NotFoundHandler())
}
