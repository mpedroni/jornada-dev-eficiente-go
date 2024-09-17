package book

import (
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
	fmt.Println("create book handler")
}
