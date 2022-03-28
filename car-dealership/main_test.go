package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/zopsmart/GoLang-Interns-2022/models"
	"github.com/zopsmart/GoLang-Interns-2022/types/brand"
	"github.com/zopsmart/GoLang-Interns-2022/types/fuel"

	"github.com/stretchr/testify/assert"
)

// Test_Main function runs the integration testing for routers
func Test_Main(t *testing.T) {
	go main()
	time.Sleep(time.Second * 1)

	data := models.Car{Name: "ABC", Year: 2020, Brand: brand.Mercedes,
		Fuel: fuel.Petrol, Engine: models.Engine{Displacement: 2200, NCylinders: 4}}
	id := "b58b91c4-735f-4e35-a6b9-8a2dc491374b"

	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	endpoint := "http://localhost:8080/cars"

	cases := []struct {
		desc       string
		method     string
		url        string
		body       io.Reader
		statusCode int
	}{
		{"create new car", "POST", endpoint, bytes.NewReader(body), http.StatusCreated},
		{"get by brand", "GET", endpoint + "?brand=Mercedes&include=true", http.NoBody, http.StatusOK},
		{"get by id", "GET", endpoint + "/" + id, http.NoBody, http.StatusOK},
		{"update existing car", "PUT", endpoint + "/" + data.ID.String(), bytes.NewReader(body), http.StatusOK},
		{"delete existing car", "DELETE", endpoint + "/" + data.ID.String(), http.NoBody, http.StatusNoContent},
		{"failure case", "HEAD", endpoint, http.NoBody, http.StatusMethodNotAllowed},
	}

	for _, tc := range cases {
		req, err := http.NewRequest(tc.method, tc.url, tc.body)
		if err != nil {
			t.Errorf("error in generating new request")

			return
		}

		req.Header.Set("Api-Key", "srijan-zs")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Errorf("error in getting response")

			return
		}

		err = res.Body.Close()
		if err != nil {
			t.Errorf("error in closing response body")

			return
		}

		assert.Equal(t, tc.statusCode, res.StatusCode, tc.desc)
	}
}
