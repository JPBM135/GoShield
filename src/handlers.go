package main

import (
	"net/http"

	"jpbm135.go-shield/src/handlers"
	"jpbm135.go-shield/src/utils"
)

func Handlers() {
	http.HandleFunc("/hash", utils.HandlerHelper(map[string]utils.HandlerFunc{
		http.MethodPost: handlers.POSTHashHandler,
	}))

	http.HandleFunc("/", utils.HandlerHelper(map[string]utils.HandlerFunc{
		http.MethodGet: handlers.RootHandler,
	}))
}
