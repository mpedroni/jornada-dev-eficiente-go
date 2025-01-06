package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "timestamp": time.Now().UTC(), "message": "Bad request"})
}

func BadRequestWithDetail(ctx *gin.Context, err error, details any) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "timestamp": time.Now().UTC(), "message": "Bad request", "detail": details})
}

func InternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong on our side", "timestamp": time.Now().UTC(), "message": "Internal server error"})
}

func Conflict(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusConflict, gin.H{"error": err.Error(), "timestamp": time.Now().UTC(), "message": "Conflict"})
}

func Created(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}
