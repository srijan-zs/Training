package main

import (
	"log"
	"net/http"

	"hello-server/handlers"
)

// main function for the hello-server, starts the server on the given address and operates accordingly
func main() {
	http.HandleFunc("/hello", handlers.Hello)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
