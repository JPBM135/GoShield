package utils

import (
	"log"
	"net/http"
)

type HandlerFunc = func(writer http.ResponseWriter, request *http.Request)

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	var message string = "Method " + request.Method + " not allowed on path " + request.URL.Path

	log.Println(message)
	WriteError(writer, message, http.StatusMethodNotAllowed)
}

func HandlerHelper(input map[string]HandlerFunc) func(http.ResponseWriter, *http.Request) {
	defaultHandlers := map[string]HandlerFunc{
		http.MethodGet:     defaultHandler,
		http.MethodPost:    defaultHandler,
		http.MethodPut:     defaultHandler,
		http.MethodDelete:  defaultHandler,
		http.MethodHead:    defaultHandler,
		http.MethodOptions: defaultHandler,
		http.MethodPatch:   defaultHandler,
	}

	for method, handler := range input {
		defaultHandlers[method] = handler
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		handler, found := defaultHandlers[request.Method]

		if !found {
			handler = defaultHandler
		}

		handler(writer, request)
	}
}
