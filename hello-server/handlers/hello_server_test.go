package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHello function contains the test cases which tests the Hello function of handler package
func TestHello(t *testing.T) {
	cases := []struct {
		desc       string
		method     string
		url        string
		output     string
		statusCode int
	}{
		{"No name parameter", "GET", "/hello", "hello", http.StatusOK},
		{"Empty name parameter", "GET", "/hello?name=", "Missing name parameter!", http.StatusBadRequest},
		{"Invalid method", "PUT", "/hello", "Only get method is allowed!", http.StatusMethodNotAllowed},
		{"Success Case", "GET", "/hello?name=Srijan", "hello Srijan", http.StatusOK},
		{"multiple name parameter", "GET", "/hello?name=Srijan&name=Shubham", "hello Srijan", http.StatusOK},
	}

	for i, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, nil)
		w := httptest.NewRecorder()

		Hello(w, req)
		resp := w.Result()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error in reading body: %v", err)
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.statusCode, resp.StatusCode)
		}

		if string(body) != tc.output {
			t.Errorf("TEST[%d], failed. %s\nExpected: %v\nGot: %v", i+1, tc.desc, tc.output, string(body))
		}
	}
}

// BenchmarkHello function runs the benchmark tests for the Hello function of handler package
func BenchmarkHello(b *testing.B) {
	cases := []struct {
		desc       string
		method     string
		url        string
		output     string
		statusCode int
	}{
		{"No name parameter", "GET", "/hello", "hello", http.StatusOK},
		{"Empty name parameter", "GET", "/hello?name=", "Missing name parameter!", http.StatusBadRequest},
		{"Invalid method", "PUT", "/hello", "Only get method is allowed!", http.StatusMethodNotAllowed},
		{"Success Case", "GET", "/hello?name=Srijan", "hello Srijan", http.StatusOK},
		{"multiple name parameter", "GET", "/hello?name=Srijan&name=Shubham", "hello Srijan", http.StatusOK},
	}

	for _, tc := range cases {
		for i := 0; i < b.N; i++ {
			req := httptest.NewRequest(tc.method, tc.url, nil)
			w := httptest.NewRecorder()

			Hello(w, req)
			resp := w.Result()

			_, _ = io.ReadAll(resp.Body)
		}
	}
}
