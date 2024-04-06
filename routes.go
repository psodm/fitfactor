package main

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/healthz", handleHealthzPlease())
	mux.Handle("/", http.NotFoundHandler())
}
