package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /authors", func(w http.ResponseWriter, r *http.Request) {
		var body CreateAuthorRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "unable to decode the request body")
			return
		}

		fmt.Println(body)
	})

	http.ListenAndServe(":8080", mux)
}
