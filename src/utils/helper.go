package utils

import (
	"net/http"
)

type HandlerHelperInputFunction = func(writer http.ResponseWriter, request *http.Request)

type HandlerHelperInput struct {
	GET    HandlerHelperInputFunction
	POST   HandlerHelperInputFunction
	PUT    HandlerHelperInputFunction
	DELETE HandlerHelperInputFunction
	HEAD   HandlerHelperInputFunction
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	var message string = "Method " + request.Method + " not allowed on path " + request.URL.Path

	WriteError(writer, message, http.StatusMethodNotAllowed)
}

func HandlerHelper(path string, input HandlerHelperInput) func(http.ResponseWriter, *http.Request) {
	if input.GET == nil {
		input.GET = defaultHandler
	}

	if input.POST == nil {
		input.POST = defaultHandler
	}

	if input.PUT == nil {
		input.PUT = defaultHandler
	}

	if input.DELETE == nil {
		input.DELETE = defaultHandler
	}

	if input.HEAD == nil {
		input.HEAD = defaultHandler
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			input.GET(writer, request)
		case "POST":
			input.POST(writer, request)
		case "PUT":
			input.PUT(writer, request)
		case "DELETE":
			input.DELETE(writer, request)
		case "HEAD":
			input.HEAD(writer, request)
		default:
			defaultHandler(writer, request)
		}
	}
}
