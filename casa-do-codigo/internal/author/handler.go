package author

import (
	"casadocodigo/internal/rest"

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

	rest.ValidateStruct(c, req)

	h.authors.Create(c.Request.Context(), req.Name, req.Email, req.Description)
}
