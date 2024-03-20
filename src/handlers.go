package main

import (
	"net/http"
	"os"

	"jpbm135.go-shield/src/handlers"
	"jpbm135.go-shield/src/utils"
)

var INFO_STRUC = map[string]interface{}{
	"author":    "jpbm135",
	"version":   "1.0.0",
	"language":  "Go",
	"goVersion": os.Environ(),
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	utils.WriteError(writer, "Not found", http.StatusNotFound)
}

func Handlers() {
	http.HandleFunc("/hash", utils.HandlerHelper("/hash", utils.HandlerHelperInput{
		POST: handlers.POSTHashHandler,
	}))

	http.HandleFunc("/", notFoundHandler)
}
