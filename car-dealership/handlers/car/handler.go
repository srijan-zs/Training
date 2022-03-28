package car

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/handlers"
	"github.com/zopsmart/GoLang-Interns-2022/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type handler struct {
	service handlers.Service
}

func New(svc handlers.Service) handler { // nolint:revive // handler is not exported so factory function is the only way to use the methods
	return handler{service: svc}
}

// Create method generates a request to create a car in handler layer
func (h handler) Create(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	car, err := getCar(req)

	if err != nil {
		setResponse(w, req.Method, car, err)

		return
	}

	data, err := h.service.Create(ctx, car)
	setResponse(w, req.Method, data, err)
}

// GetAll method generates a request to extract cars using filter of brand and engine
func (h handler) GetAll(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	params := mux.Vars(req)

	brand := strings.TrimSpace(params["brand"])
	include := strings.TrimSpace(params["include"])

	engine := false
	if strings.EqualFold(include, "true") {
		engine = true
	}

	filter := filters.Car{
		Brand:         brand,
		IncludeEngine: engine,
	}

	data, err := h.service.GetAll(ctx, filter)
	setResponse(w, req.Method, data, err)
}

// GetByID method generates a request to extract car of a given id
func (h handler) GetByID(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getID(req)
	if err != nil {
		setResponse(w, req.Method, nil, err)

		return
	}

	data, err := h.service.GetByID(ctx, id)
	setResponse(w, req.Method, data, err)
}

// Update method generates a request to update the existing car details
func (h handler) Update(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getID(req)
	if err != nil {
		setResponse(w, req.Method, nil, err)

		return
	}

	car, err := getCar(req)
	if err != nil {
		setResponse(w, req.Method, car, err)

		return
	}

	car.ID = id

	car, err = h.service.Update(ctx, car.ID, car)
	setResponse(w, req.Method, car, err)
}

// Delete method generates a request to delete the existing car of a given id
func (h handler) Delete(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getID(req)
	if err != nil {
		setResponse(w, req.Method, nil, err)

		return
	}

	err = h.service.Delete(ctx, id)
	setResponse(w, req.Method, nil, err)
}

// getCar function reads the body of request and unmarshal it to create a car model
func getCar(req *http.Request) (*models.Car, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.InvalidParam{Params: []string{"body"}}
	}

	var car models.Car

	err = json.Unmarshal(body, &car)
	if err != nil {
		return nil, errors.InvalidParam{Params: []string{"body"}}
	}

	return &car, nil
}

// getID function reads the id from the request given as a query parameter
func getID(req *http.Request) (uuid.UUID, error) {
	params := mux.Vars(req)
	id := strings.TrimSpace(params["id"])

	if id == "" {
		return uuid.Nil, errors.MissingParam{Params: []string{"id"}}
	}

	ID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.InvalidParam{Params: []string{"id"}}
	}

	return ID, nil
}
