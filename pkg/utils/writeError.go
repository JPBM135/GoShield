package utils

import "net/http"

func WriteError(writer http.ResponseWriter, message string, code int) {
	writer.WriteHeader(code)
	writer.Write([]byte(message))
}
