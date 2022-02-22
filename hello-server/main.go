package main

import (
	"log"
	"net/http"

	"github.com/srijan-zs/hello-server/handler"
)

// main function for the hello-server, starts the server on the given address and operates accordingly
func main() {
	http.HandleFunc("/hello", handler.Hello)

	log.Fatalln(http.ListenAndServe(":8000", nil))
}
