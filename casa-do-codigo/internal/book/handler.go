package book

import (
	"casadocodigo/internal/rest"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	Create(c *gin.Context)
}

type bookHandler struct{}

func NewBookHandler() BookHandler {
	return &bookHandler{}
}

func (h *bookHandler) Create(c *gin.Context) {
	req := &CreateBookRequest{}

	if !rest.ValidateStruct(c, req) {
		return
	}

	fmt.Println("create book handler", req)
}
