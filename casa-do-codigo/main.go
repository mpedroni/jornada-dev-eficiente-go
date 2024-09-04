package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func (c *CreateAuthorRequest) Validate() error {
	return v.ValidateStruct(c,
		v.Field(&c.Name, v.Required),
		v.Field(&c.Email, v.Required, is.Email),
		v.Field(&c.Description, v.Required, v.Length(0, 400)),
	)
}

type HttpErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Detail    string    `json:"detail,omitempty"`
	Details   []string  `json:"details,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /authors", func(w http.ResponseWriter, r *http.Request) {
		body := &CreateAuthorRequest{}
		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(HttpErrorResponse{
				Timestamp: time.Now().UTC(),
				Message:   http.StatusText(http.StatusBadRequest),
				Detail:    "Invalid request body",
			})
			return
		}

		if err := body.Validate(); err != nil {
			ee := err.(v.Errors)
			details := make([]string, 0)
			for k, v := range ee {
				details = append(details, k+": "+v.Error())
			}
			resp := HttpErrorResponse{
				Timestamp: time.Now().UTC(),
				Message:   http.StatusText(http.StatusBadRequest),
				Detail:    "One or more invalid fields. Fix it and try again.",
				Details:   details,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)

			return
		}
	})

	http.ListenAndServe(":8080", mux)
}
