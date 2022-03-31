package book

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/srijan-zs/Training/library-management/errors"
	"github.com/srijan-zs/Training/library-management/handlers"
	"github.com/srijan-zs/Training/library-management/models"
	"io"
	"net/http"
	"strings"
)

type handler struct {
	service handlers.Service
}

func New(svc handlers.Service) handler {
	return handler{service: svc}
}

func (h handler) Create(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	book, err := getBook(req)

	if err != nil {
		setResponse(w, req.Method, book, err)
	}

	data, err := h.service.Create(ctx, book)
	setResponse(w, req.Method, data, err)
}

func (h handler) Get(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getID(req)
	if err != nil {
		setResponse(w, req.Method, nil, err)

		return
	}

	data, err := h.service.Get(ctx, id)
	setResponse(w, req.Method, data, err)
}

func (h handler) Update(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	id, err := getID(req)
	if err != nil {
		setResponse(w, req.Method, nil, err)

		return
	}

	book, err := getBook(req)
	if err != nil {
		setResponse(w, req.Method, book, err)

		return
	}

	book.ID = id

	book, err = h.service.Update(ctx, book.ID, book)
	setResponse(w, req.Method, book, err)
}

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

func getBook(req *http.Request) (*models.Book, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.InvalidParam{Params: []string{"body"}}
	}

	var book models.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		return nil, errors.InvalidParam{Params: []string{"body"}}
	}

	return &book, nil
}

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
