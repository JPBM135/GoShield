package main

import (
	"net/http"

	"jpbm135.go-shield/src/handlers"
	"jpbm135.go-shield/src/utils"
)

var INFO_STRUCTURE = map[string]interface{}{
	"author":    "jpbm135",
	"version":   "1.0.0",
	"language":  "Go",
	"goVersion": "1.20",
}

func notFoundHandler(writer http.ResponseWriter) {
	utils.WriteError(writer, "Not found", http.StatusNotFound)
}

func Handlers() {
	http.HandleFunc("/hash", utils.HandlerHelper("/hash", utils.HandlerHelperInput{
		POST: handlers.POSTHashHandler,
	}))

	http.HandleFunc("/", utils.HandlerHelper("/", utils.HandlerHelperInput{
		GET: func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Path != "/" {
				notFoundHandler(writer)
				return
			}
			utils.WriteJSON(writer, INFO_STRUCTURE)
		},
	}))
}
