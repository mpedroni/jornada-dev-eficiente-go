package handler

import (
	"cdc-v2/internal/author/domain"
	"cdc-v2/pkg/rest"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	validator *validator.Validate
	repo      domain.AuthorRepository
}

func New(
	validator *validator.Validate,
	repository domain.AuthorRepository,
) *handler {
	return &handler{
		validator: validator,
		repo:      repository,
	}
}

func (h *handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/authors", h.CreateAuthor)
}

func (h *handler) CreateAuthor(ctx *gin.Context) {
	var req CreateAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rest.BadRequest(ctx, err)
		return
	}

	if detail := rest.ValidateJSON(req); detail != nil {
		rest.BadRequestWithDetail(ctx, errors.New("One or more fields are invalid or missing"), detail)
		return
	}

	if author, err := h.repo.GetByEmail(ctx, req.Email); err != nil && !errors.Is(err, domain.ErrAuthorNotFound) {
		rest.InternalServerError(ctx)
		return
	} else if author != nil {
		rest.Conflict(ctx, errors.New("Already exists an author with the given email"))
		return
	}

	author := domain.NewAuthor(req.Name, req.Email, req.Description)

	if err := h.repo.Create(ctx, author); err != nil {
		rest.InternalServerError(ctx)
		return
	}

	rest.Created(ctx)
}
