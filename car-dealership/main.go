package main

import (
	"log"
	"net/http"

	"github.com/zopsmart/GoLang-Interns-2022/drivers"
	handlers "github.com/zopsmart/GoLang-Interns-2022/handlers/car"
	"github.com/zopsmart/GoLang-Interns-2022/middlewares"
	services "github.com/zopsmart/GoLang-Interns-2022/services/car"
	"github.com/zopsmart/GoLang-Interns-2022/stores/car"
	"github.com/zopsmart/GoLang-Interns-2022/stores/engine"

	"github.com/gorilla/mux"
)

func main() {
	db, err := drivers.Connection()
	if err != nil {
		log.Printf("Database can't be connected: %v", err)
	}

	log.Printf("Database connected")

	carStore := car.New(db)
	engineStore := engine.New(db)

	service := services.New(carStore, engineStore)
	handler := handlers.New(service)

	r := mux.NewRouter()

	r.HandleFunc("/cars", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/cars", handler.GetAll).Queries("brand", "{brand}", "include", "{include}").Methods(http.MethodGet)
	r.HandleFunc("/cars/{id}", handler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id}", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/cars/{id}", handler.Delete).Methods(http.MethodDelete)

	r.Use(middlewares.Authentication)

	log.Fatal(http.ListenAndServe(":8080", r))
}
