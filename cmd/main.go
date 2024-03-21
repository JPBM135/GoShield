package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"jpbm135.go-shield/pkg/router"
)

func main() {
	fmt.Println("starting GoShield")

	shutdown := make(chan os.Signal, 1)
	serverErrors := make(chan error, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	r := router.New()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		fmt.Println("Starting server on port 8080")
		serverErrors <- s.ListenAndServe()
	}()
	select {
	case err := <-serverErrors:
		fmt.Println("server error:", err)

	case sig := <-shutdown:
		fmt.Println("shutdown initiated:", sig)

		// Give requests time to finish
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		// Ask the server to shutdown gracefully
		if err := s.Shutdown(ctx); err != nil {
			fmt.Println("graceful shutdown failed:", err)
			err = s.Close()
		}
	}
}
