-- name: CreateUser :one
INSERT INTO users (login, password, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id;