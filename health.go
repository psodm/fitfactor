package main

import "net/http"

type Health struct {
	Message string `json:"message"`
}

func handleHealthzPlease() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			encode(w, r, http.StatusOK, Health{"Hello World"})
		},
	)
}
