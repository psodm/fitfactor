package main

import (
	"context"
	"net/http"
)

func NewServer(ctx context.Context, userStore *UserStore) http.Handler {
	mux := http.NewServeMux()
	addRoutes(ctx, mux, userStore)
	var handler http.Handler = mux
	return handler
}
