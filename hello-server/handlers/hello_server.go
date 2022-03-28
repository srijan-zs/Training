package handlers

import (
	"net/http"
)

// Hello function is the handler function for the hello-server and checks URL for different cases to write the desired outcome
func Hello(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Only get method is allowed!"))

		return
	}

	params := req.URL.Query()

	if len(params) > 0 {
		name := params.Get("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing name parameter!"))

			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello " + name))

		return
	}

	_, _ = w.Write([]byte("hello"))
}
