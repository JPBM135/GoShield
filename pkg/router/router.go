package router

import (
	"net/http"

	"jpbm135.go-shield/pkg/handlers"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", handlers.RootHandler)
	r.HandleFunc("POST /hash", handlers.POSTHashHandler)

	return r
}
