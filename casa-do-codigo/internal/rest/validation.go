package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(c *gin.Context, obj interface{}) bool {
	payload := gin.H{
		"timestamp": time.Now().UTC(),
		"message":   "Bad Request",
		"status":    http.StatusBadRequest,
	}

	if err := c.ShouldBindJSON(obj); err != nil {
		if ee, ok := err.(validator.ValidationErrors); ok {
			var messages []string
			for _, validationErr := range ee {
				field := validationErr.Field()
				tag := validationErr.Tag()

				var message string
				switch tag {
				case "required":
					message = field + " is required"
				case "tags":
					message = field + " cannot have more than 5 tags"
				case "duplicated":
					message = field + " has duplicated values"
				case "not_blank":
					message = field + " cannot be empty or contain only spaces"
				case "max":
					message = field + " must not exceed " + validationErr.Param() + " characters"
				case "min":
					message = field + " must have at least " + validationErr.Param() + " characters"
				default:
					message = field + " is invalid"
				}

				messages = append(messages, message)
			}

			payload["detail"] = "Invalid request. Verify the fields and try again"
			payload["details"] = messages

			c.JSON(http.StatusBadRequest, payload)
			return false
		}

		payload["detail"] = "Unable to parse: " + err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return false
	}

	return true
}
