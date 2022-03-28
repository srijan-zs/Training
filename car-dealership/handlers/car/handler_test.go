package car

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/handlers"
	"github.com/zopsmart/GoLang-Interns-2022/models"
	"github.com/zopsmart/GoLang-Interns-2022/types/brand"
	"github.com/zopsmart/GoLang-Interns-2022/types/fuel"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// initializeTest generates the mocks for creating mock calls to test the defined methods
func initializeTest(t *testing.T, method string, body io.Reader) (
	*handlers.MockService, handler, *http.Request, *httptest.ResponseRecorder) {
	ctrl := gomock.NewController(t)

	mockService := handlers.NewMockService(ctrl)
	h := New(mockService)

	req := httptest.NewRequest(method, "/cars", body)

	w := httptest.NewRecorder()

	return mockService, h, req, w
}

// car is created as a model for testing functions to use as a sample data
var car = models.Car{ // nolint:gochecknoglobals // variable declared as global to reduce redundancy in tests
	ID:     uuid.MustParse("36da908c-5c03-42b9-bec4-20a66446e322"),
	Name:   "ABC",
	Year:   2020,
	Brand:  brand.Mercedes,
	Fuel:   fuel.Petrol,
	Engine: models.Engine{Displacement: 2000, NCylinders: 4},
}

// TestHandler_Create function checks the test cases for Create method of handler layer
func TestHandler_Create(t *testing.T) {
	body := []byte(`{"id":"36da908c-5c03-42b9-bec4-20a66446e322","name":"ABC","year":2020,"brand":"mercedes","fuel":"petrol",
		"engine":{"displacement":2000,"nCylinder":4}}`)

	cases := []struct {
		desc       string
		mockOutput *models.Car
		mockErr    error
		resp       *models.Car
		statusCode int
	}{
		{"success", &car, nil, &car, http.StatusCreated},
		{"entity already exists", nil, errors2.EntityAlreadyExists{}, nil, http.StatusOK},
		{"internal server error", nil, errors2.DB{}, nil, http.StatusInternalServerError},
	}

	for _, tc := range cases {
		mockService, h, r, w := initializeTest(t, http.MethodPost, bytes.NewReader(body))

		mockService.EXPECT().Create(gomock.Any(), &car).Return(tc.mockOutput, tc.mockErr)

		h.Create(w, r)
		resp := w.Result()

		body, err := getResponseBody(resp)
		if err != nil {
			t.Errorf("error in reading body: %v", err)
		}

		output := getOutput(t, body)

		assert.Equal(t, tc.statusCode, resp.StatusCode, tc.desc)

		assert.Equal(t, tc.resp, output, tc.desc)
	}
}

// mockBody is the nil struct defined to generate invalid body error
type mockBody struct{}

func (m mockBody) Read(p []byte) (n int, err error) {
	return 0, errors2.InvalidParam{}
}

// Test_CreateInvalidBody function checks the test cases for invalid body
func Test_CreateInvalidBody(t *testing.T) {
	_, h, r, w := initializeTest(t, http.MethodPost, mockBody{})

	h.Create(w, r)
	resp := w.Result()

	body, err := getResponseBody(resp)
	if err != nil {
		t.Errorf("error in reading body: %v", err)
	}

	output := getOutput(t, body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	if output != nil {
		t.Errorf("[TEST] Failed. Desc : invalid body\nGot: %v\nExpected: %v", output, resp.Body)
	}
}

// TestHandler_GetAll function checks the test cases for read all cars with specified brand and engine
func TestHandler_GetAll(t *testing.T) {
	id := uuid.New()

	withEngine := []models.Car{
		{
			ID:     id,
			Name:   "ABC",
			Year:   2020,
			Brand:  brand.Mercedes,
			Fuel:   fuel.Petrol,
			Engine: models.Engine{Displacement: 2000, NCylinders: 4},
		},
	}

	withoutEngine := []models.Car{
		{
			ID:    id,
			Name:  "ABC",
			Year:  2020,
			Brand: brand.Mercedes,
			Fuel:  fuel.Petrol,
		},
	}

	cases := []struct {
		desc       string
		filter     filters.Car
		mockOutput []models.Car
		mockErr    error
		statusCode int
	}{
		{"get all cars with engine", filters.Car{Brand: "mercedes", IncludeEngine: true}, withEngine, nil, http.StatusOK},
		{"get all cars without engine", filters.Car{Brand: "mercedes", IncludeEngine: false}, withoutEngine, nil, http.StatusOK},
		{"invalid parameter", filters.Car{Brand: "abc"}, nil, errors2.InvalidParam{Params: []string{"brand"}}, http.StatusBadRequest},
		{"internal server error", filters.Car{Brand: "bmw"}, nil, errors2.DB{}, http.StatusInternalServerError},
	}

	for _, tc := range cases {
		mockService, h, r, w := initializeTest(t, http.MethodGet, http.NoBody)

		req := mux.SetURLVars(r, map[string]string{"brand": tc.filter.Brand, "include": strconv.FormatBool(tc.filter.IncludeEngine)})

		mockService.EXPECT().GetAll(gomock.Any(), tc.filter).Return(tc.mockOutput, tc.mockErr)

		h.GetAll(w, req)
		resp := w.Result()

		_, err := getResponseBody(resp)
		if err != nil {
			t.Errorf("error in reading body: %v", err)
		}

		assert.Equal(t, tc.statusCode, resp.StatusCode, tc.desc)
	}
}

// TestHandler_GetByID function checks the test cases for extracting the car of a specific id
func TestHandler_GetByID(t *testing.T) {
	id := uuid.New()
	car.ID = id

	cases := []struct {
		desc       string
		mockOutput models.Car
		mockErr    error
		statusCode int
	}{
		{"success", car, nil, http.StatusOK},
		{"entity does not exists", models.Car{}, errors2.EntityNotFound{}, http.StatusNotFound},
		{"internal server error", models.Car{}, errors2.DB{}, http.StatusInternalServerError},
	}

	for _, tc := range cases {
		mockService, h, r, w := initializeTest(t, http.MethodGet, http.NoBody)
		req := mux.SetURLVars(r, map[string]string{"id": id.String()})

		mockService.EXPECT().GetByID(gomock.Any(), id).Return(tc.mockOutput, tc.mockErr)

		h.GetByID(w, req)
		resp := w.Result()

		_, err := getResponseBody(resp)
		if err != nil {
			t.Errorf("error in reading body: %v", err)
		}

		assert.Equal(t, tc.statusCode, resp.StatusCode, tc.desc)
	}
}

// Test_GetByIDInvalidID function checks the case for invalid id given
func Test_GetByIDInvalidID(t *testing.T) {
	_, h, r, w := initializeTest(t, http.MethodGet, http.NoBody)

	h.GetByID(w, r)
	resp := w.Result()

	body, err := getResponseBody(resp)
	if err != nil {
		t.Errorf("error in reading body : %v", err)
	}

	output := getOutput(t, body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	if output != nil {
		t.Errorf("[TEST] Failed. Desc : invalid body\nGot: %v\nExpected: %v", string(body), resp.Body)
	}
}

// TestHandler_Update function checks the test cases for update method of handler layer
func TestHandler_Update(t *testing.T) {
	body := []byte(`{"id":"4aa70ec2-ffb7-4054-a388-6db5d840807b","name":"XYZ","year":2019,"brand":"ferrari","fuel":"petrol",
		"engine":{"displacement":2200,"nCylinder":4}}`)

	var input = models.Car{
		ID:     uuid.MustParse("4aa70ec2-ffb7-4054-a388-6db5d840807b"),
		Name:   "XYZ",
		Year:   2019,
		Brand:  brand.Ferrari,
		Fuel:   fuel.Petrol,
		Engine: models.Engine{Displacement: 2200, NCylinders: 4},
	}

	cases := []struct {
		desc       string
		mockErr    error
		resp       *models.Car
		statusCode int
	}{
		{"success", nil, &input, http.StatusOK},
		{"entity not found", errors2.EntityNotFound{}, nil, http.StatusNotFound},
		{"internal server error", errors2.DB{}, nil, http.StatusInternalServerError},
	}

	for _, tc := range cases {
		mockService, h, r, w := initializeTest(t, http.MethodPut, bytes.NewReader(body))

		req := mux.SetURLVars(r, map[string]string{"id": "4aa70ec2-ffb7-4054-a388-6db5d840807b"})

		mockService.EXPECT().Update(gomock.Any(), uuid.MustParse("4aa70ec2-ffb7-4054-a388-6db5d840807b"), &input).Return(tc.resp, tc.mockErr)

		h.Update(w, req)
		resp := w.Result()

		data, err := getResponseBody(resp)
		if err != nil {
			t.Errorf("error in reading body: %v", err)
		}

		output := getOutput(t, data)

		assert.Equal(t, tc.statusCode, resp.StatusCode, tc.desc)

		assert.Equal(t, tc.resp, output, tc.desc)
	}
}

// Test_UpdateInvalidID function checks the case for invalid id given
func Test_UpdateInvalidID(t *testing.T) {
	_, h, r, w := initializeTest(t, http.MethodPut, http.NoBody)

	req := mux.SetURLVars(r, map[string]string{"id": "36da908c-5c03-42b9-bec4-20a66446e322"})

	h.Update(w, req)
	resp := w.Result()

	body, err := getResponseBody(resp)
	if err != nil {
		t.Errorf("error in reading body: %v", err)
	}

	output := getOutput(t, body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	if output != nil {
		t.Errorf("[TEST] Failed. Desc : invalid body\nGot: %v\nExpected: %v", string(body), resp.Body)
	}
}

// Test_UpdateInvalidBody function checks the test cases for invalid body
func Test_UpdateInvalidBody(t *testing.T) {
	_, h, r, w := initializeTest(t, http.MethodPut, http.NoBody)

	h.Update(w, r)
	resp := w.Result()

	body, err := getResponseBody(resp)
	if err != nil {
		t.Errorf("error in reading body: %v", err)
	}

	output := getOutput(t, body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	if output != nil {
		t.Errorf("[TEST] Failed. Desc : invalid body\nGot: %v\nExpected: %v", string(body), resp.Body)
	}
}

// TestHandler_Delete function checks the test cases for delete method of handler layer
func TestHandler_Delete(t *testing.T) {
	id := uuid.New()

	cases := []struct {
		desc       string
		mockErr    error
		statusCode int
	}{
		{"success", nil, http.StatusNoContent},
		{"entity does not exists", errors2.EntityNotFound{}, http.StatusNotFound},
		{"internal server error", errors2.DB{}, http.StatusInternalServerError},
	}

	for _, tc := range cases {
		mockService, h, r, w := initializeTest(t, http.MethodDelete, http.NoBody)

		mockService.EXPECT().Delete(gomock.Any(), id).Return(tc.mockErr)

		req := mux.SetURLVars(r, map[string]string{"id": id.String()})

		h.Delete(w, req)
		resp := w.Result()

		err := resp.Body.Close()
		if err != nil {
			t.Errorf("error in closing response: %v", err)
		}

		assert.Equal(t, tc.statusCode, resp.StatusCode, tc.desc)
	}
}

// Test_DeleteInvalidID function checks the case for invalid id given
func Test_DeleteInvalidID(t *testing.T) {
	_, h, r, w := initializeTest(t, http.MethodDelete, http.NoBody)

	h.Delete(w, r)
	resp := w.Result()

	body, err := getResponseBody(resp)
	if err != nil {
		t.Errorf("error in reading body: %v", err)
	}

	output := getOutput(t, body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	if output != nil {
		t.Errorf("[TEST] Failed. Desc : invalid body\nGot: %v\nExpected: %v", string(body), resp.Body)
	}
}

// Test_getCar function checks the test cases for getCar function of handler layer
func Test_getCar(t *testing.T) {
	body := bytes.NewReader([]byte(`{"id":"36da908c-5c03-42b9-bec4-20a66446e322","name":"ABC","year":2020,"brand":"mercedes","fuel":"petrol",
		"engine":{"displacement":2000,"nCylinder":4}}`))

	invalidBody := bytes.NewReader([]byte("invalid body"))

	cases := []struct {
		desc   string
		body   io.Reader
		output *models.Car
		err    error
	}{
		{"success", body, &car, nil},
		{"bind error", mockBody{}, nil, errors2.InvalidParam{Params: []string{"body"}}},
		{"unmarshal error", invalidBody, nil, errors2.InvalidParam{Params: []string{"body"}}},
	}

	for _, tc := range cases {
		_, _, r, _ := initializeTest(t, "", tc.body)

		_, err := getCar(r)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// Test_getID function checks the test cases for getID function of handler layer
func Test_getID(t *testing.T) {
	cases := []struct {
		desc   string
		id     string
		output uuid.UUID
		err    error
	}{
		{"id parsed success", "36da908c-5c03-42b9-bec4-20a66446e322", uuid.MustParse("36da908c-5c03-42b9-bec4-20a66446e322"), nil},
		{"empty string", "", uuid.Nil, errors2.MissingParam{Params: []string{"id"}}},
		{"invalid id", "123456", uuid.Nil, errors2.InvalidParam{Params: []string{"id"}}},
	}

	for _, tc := range cases {
		_, _, r, _ := initializeTest(t, "", nil)

		req := mux.SetURLVars(r, map[string]string{"id": tc.id})

		output, err := getID(req)

		assert.Equal(t, tc.err, err, tc.desc)

		assert.Equal(t, tc.output, output, tc.desc)
	}
}

// Test_writeResponseBodyMarshalError function checks the case for marshal error
func Test_writeResponseBodyMarshalError(t *testing.T) {
	data := complex(1, 1)

	w := httptest.NewRecorder()
	expectedStatusCode := http.StatusInternalServerError

	writeResponseBody(w, http.StatusOK, data)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, expectedStatusCode, resp.StatusCode)
}

// getResponseBody function reads the response body
func getResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

// getOutput function unmarshal the data to type car model
func getOutput(t *testing.T, respBody []byte) *models.Car {
	var output *models.Car

	if len(respBody) != 0 {
		output = &models.Car{}

		err := json.Unmarshal(respBody, output)
		if err != nil {
			t.Error(err)
		}
	}

	return output
}
