package author

import (
	"casadocodigo/internal/rest"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthorHandler interface {
	Create(c *gin.Context)
}

type authorHandler struct {
	authors AuthorService
}

func NewAuthorHandler(as AuthorService) AuthorHandler {
	return &authorHandler{
		authors: as,
	}
}

func (h *authorHandler) Create(c *gin.Context) {
	req := &CreateAuthorRequest{}

	if !rest.ValidateStruct(c, req) {
		return
	}

	_, err := h.authors.Create(c.Request.Context(), req.Name, req.Email, req.Description)
	if err != nil {
		if err == ErrEmailAlreadyExists {
			rest.BadRequest(c, fmt.Sprintf("An author with email %s already exists", req.Email))
			return
		}

		log.Printf("while creating author: %v\n", err)
		rest.InternalServerError(c)
		return
	}

	rest.OK(c)
}
