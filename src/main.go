package main

import (
	"log"
	"net/http"
)

func main() {
	Handlers()
	log.Default().Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
