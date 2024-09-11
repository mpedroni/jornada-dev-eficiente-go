package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Detail    string    `json:"detail,omitempty"`
	Details   []string  `json:"details,omitempty"`
}

func httpErrorResponse(c *gin.Context, status int, message, detail string, details ...string) {
	c.JSON(status, HttpErrorResponse{
		Timestamp: time.Now().UTC(),
		Message:   message,
		Detail:    detail,
		Details:   details,
	})
}

func InternalServerError(c *gin.Context) {
	status := http.StatusInternalServerError

	httpErrorResponse(
		c,
		status,
		http.StatusText(status),
		"An unexpected error occurred. Please try again later",
	)
}

func BadRequest(c *gin.Context, detail string, details ...string) {
	status := http.StatusBadRequest

	httpErrorResponse(
		c,
		status,
		http.StatusText(status),
		detail,
		details...,
	)
}

func OK(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.Status(http.StatusOK)
		return
	}
	c.JSON(http.StatusOK, data[0])
}
