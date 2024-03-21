package handlers

import (
	"net/http"

	"jpbm135.go-shield/src/utils"
)

var INFO_STRUCTURE = map[string]interface{}{
	"author":    "jpbm135",
	"version":   "1.0.0",
	"goVersion": "1.20",
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		NotFoundHandler(writer)
		return
	}

	utils.WriteJSON(writer, INFO_STRUCTURE)
}
