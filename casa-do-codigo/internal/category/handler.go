package category

import (
	"casadocodigo/internal/rest"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	Create(c *gin.Context)
}

type categoryHandler struct {
	categories CategoryService
}

func NewCategoryHandler(c CategoryService) CategoryHandler {
	return &categoryHandler{
		categories: c,
	}
}

func (h *categoryHandler) Create(c *gin.Context) {
	req := &CreateCategoryRequest{}

	if !rest.ValidateStruct(c, req) {
		return
	}

	_, err := h.categories.Create(c.Request.Context(), req.Name)
	if err != nil {
		if errors.Is(err, ErrCategoryAlreadyExists) {
			rest.BadRequest(c, fmt.Sprintf("A category with name %s already exists", req.Name))
			return
		}

		rest.InternalServerError(c)
		return
	}
}
