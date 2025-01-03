package handler

import (
	"cdc-v2/internal/author/domain"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "timestamp": time.Now().Unix(), "message": "Bad request"})
		return
	}

	if err := h.validator.Struct(req); err != nil {

		var errors map[string][]string = make(map[string][]string)
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

			errors[a.Tag.Get("json")] = append(errors[err.Field()], message)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors, "timestamp": time.Now().Unix(), "message": "Bad request"})
		return
	}

	if author, err := h.repo.GetByEmail(ctx, req.Email); err != nil && !errors.Is(err, domain.ErrAuthorNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "timestamp": time.Now().Unix(), "message": "Internal server error"})
		return
	} else if author != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "author already exists", "timestamp": time.Now().Unix(), "message": "Conflict"})
		return
	}

	author := domain.NewAuthor(req.Name, req.Email, req.Description)

	if err := h.repo.Create(ctx, author); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "timestamp": time.Now().Unix(), "message": "Internal server error"})
		return
	}

	ctx.Status(http.StatusCreated)
}
