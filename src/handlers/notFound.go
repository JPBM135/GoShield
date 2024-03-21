package handlers

import (
	"net/http"

	"jpbm135.go-shield/src/utils"
)

func NotFoundHandler(writer http.ResponseWriter) {
	utils.WriteError(writer, "Not found", http.StatusNotFound)
}
