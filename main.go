package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	// _ "https://github.com/lib/pq"
)

var (
	getenv = func(key string) string {
		switch key {
		case "ENV":
			return "development"
		case "HOST":
			return "127.0.0.1"
		case "PORT":
			return "6969"
		case "SECRET-KEY":
			return "supersecretkey"
		default:
			return ""
		}
	}
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, getenv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	db, err := Connect("postgres://postgres:password@localhost:5432/fitfactor?sslmode=disable", 25, 25, 15)
	if err != nil {
		return err
	}

	userStore := UserStore{db: db}
	srv := NewServer(ctx, &userStore)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", getenv("HOST"), getenv("PORT")),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
