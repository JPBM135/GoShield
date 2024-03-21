package handlers

import (
	"log"
	"net/http"

	"jpbm135.go-shield/pkg/utils"
)

var INFO_STRUCTURE = map[string]interface{}{
	"author":    "jpbm135",
	"version":   "1.0.0",
	"goVersion": "1.20",
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {

	utils.WriteJSON(writer, INFO_STRUCTURE)
}

func NotFoundHandler(writer http.ResponseWriter) {
	// panic("unimplemented")
	log.Fatalf("this is unimplemented \n")
}
