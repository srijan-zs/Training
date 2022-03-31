package book

import (
	"encoding/json"
	"github.com/srijan-zs/Training/library-management/errors"
	"net/http"
)

func setResponse(w http.ResponseWriter, method string, data interface{}, err error) {
	switch err.(type) {
	case errors.EntityAlreadyExists:
		w.WriteHeader(http.StatusOK)
	case errors.MissingParam, errors.InvalidParam:
		w.WriteHeader(http.StatusBadRequest)
	case errors.EntityNotFound:
		w.WriteHeader(http.StatusNotFound)
	case nil:
		writeSuccessResponse(method, w, data)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeSuccessResponse(method string, w http.ResponseWriter, data interface{}) {
	switch method {
	case http.MethodPost:
		writeResponseBody(w, http.StatusCreated, data)
	case http.MethodGet:
		writeResponseBody(w, http.StatusOK, data)
	case http.MethodPut:
		writeResponseBody(w, http.StatusOK, data)
	case http.MethodDelete:
		writeResponseBody(w, http.StatusNoContent, data)
	}
}

func writeResponseBody(w http.ResponseWriter, statusCode int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(statusCode)

	_, err = w.Write(resp)
	if err != nil {
		return
	}
}
