package handler

import (
	"cdc-v2/internal/author/domain"
	"cdc-v2/pkg/rest"
	"errors"
	"fmt"
	"reflect"

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

	if err := h.validator.Struct(req); err != nil {
		var ee map[string][]string = make(map[string][]string)
		for _, err := range err.(validator.ValidationErrors) {
			fld := reflect.ValueOf(req).Type()
			a, _ := fld.FieldByName(err.Field())

			var message string
			message = a.Tag.Get(err.Tag())
			if message == "" {
				switch err.Tag() {
				case "required":
					message = fmt.Sprintf("%s is required", a.Tag.Get("json"))
				case "email":
					message = fmt.Sprintf("%s must be a valid email address", a.Tag.Get("json"))
				default:
					message = fmt.Sprintf("%s should be %s", a.Tag.Get("json"), err.Tag())
					if err.Param() != "" {
						message += fmt.Sprintf("=%s", err.Param())
					}

				}
			}

			ee[a.Tag.Get("json")] = append(ee[err.Field()], message)
		}

		rest.BadRequestWithDetail(ctx, errors.New("One or more fields are invalid"), ee)
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
