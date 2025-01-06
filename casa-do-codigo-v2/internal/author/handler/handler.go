package handler

import (
	"cdc-v2/internal/author/domain"
	"cdc-v2/pkg/rest"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type handler struct {
	repo   domain.AuthorRepository
	logger *zap.Logger
}

func New(
	repository domain.AuthorRepository,
	logger *zap.Logger,
) *handler {
	return &handler{
		repo:   repository,
		logger: logger,
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

	author, err := h.repo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, domain.ErrAuthorNotFound) {
		h.logger.Error(err.Error())
		rest.InternalServerError(ctx)
		return
	}

	if author != nil {
		rest.Conflict(ctx, errors.New("Already exists an author with the given email"))
		return
	}

	author = domain.NewAuthor(req.Name, req.Email, req.Description)

	if err := h.repo.Create(ctx, author); err != nil {
		h.logger.Error(err.Error())
		rest.InternalServerError(ctx)
		return
	}

	rest.Created(ctx)
}
