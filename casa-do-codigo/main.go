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

func httpError(w http.ResponseWriter, status int, detail string, details ...[]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var dd []string
	if len(details) > 0 {
		dd = details[0]
	}

	json.NewEncoder(w).Encode(HttpErrorResponse{
		Timestamp: time.Now().UTC(),
		Message:   http.StatusText(status),
		Detail:    detail,
		Details:   dd,
	})
}

func HttpBadRequest(w http.ResponseWriter, detail string, details ...[]string) {
	httpError(w, http.StatusBadRequest, detail, details...)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /authors", func(w http.ResponseWriter, r *http.Request) {
		body := &CreateAuthorRequest{}
		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			HttpBadRequest(w, "Invalid request body")
			return
		}

		if err := body.Validate(); err != nil {
			ee := err.(v.Errors)
			details := make([]string, 0)
			for k, v := range ee {
				details = append(details, k+": "+v.Error())
			}

			HttpBadRequest(w, "One or more invalid fields. Fix it and try again.", details)

			return
		}
	})

	http.ListenAndServe(":8080", mux)
}
