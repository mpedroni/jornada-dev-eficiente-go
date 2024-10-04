package main

import (
	"database/sql"
	"log"
	"mercadolivre/db/melisqlc"
	"mercadolivre/internal/user"

	"github.com/gofiber/fiber/v3"

	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/desafio_meli?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := melisqlc.New(conn)

	users := user.NewUserHandler(queries)

	app.Post("/users", users.Create)

	log.Fatal(app.Listen(":3000"))
}
