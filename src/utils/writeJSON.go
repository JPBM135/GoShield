package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}
