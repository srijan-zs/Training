package main

import (
	"github.com/srijan-zs/Training/library-management/drivers"
	handlers "github.com/srijan-zs/Training/library-management/handlers/book"
	services "github.com/srijan-zs/Training/library-management/services/book"
	"github.com/srijan-zs/Training/library-management/stores/book"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := drivers.Connection()
	if err != nil {
		log.Printf("Database not connected: %v", err)
	}

	log.Printf("Database connected")

	bookStore := book.New(db)

	service := services.New(bookStore)
	handler := handlers.New(service)

	r := mux.NewRouter()

	r.HandleFunc("/books", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", handler.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", r))
}
