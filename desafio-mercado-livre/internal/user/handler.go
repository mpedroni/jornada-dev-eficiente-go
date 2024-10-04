package user

import (
	"mercadolivre/db/melisqlc"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type UserHandler struct {
	repo *melisqlc.Queries
}

func NewUserHandler(q *melisqlc.Queries) UserHandler {
	return UserHandler{
		repo: q,
	}
}

type CreateUserRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	body := new(CreateUserRequestBody)
	if err := c.Bind().JSON(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"detail":  err.Error(),
		})
	}

	user := NewUser(body.Login, body.Password)
	id, err := h.repo.CreateUser(c.Context(), melisqlc.CreateUserParams{
		Login:     user.Login,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	})
	if err != nil {
		log.Errorf("error creating user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	user.ID = int(id)

	return c.Status(fiber.StatusNoContent).Send(nil)
}
