-- name: CreateAuthor :one
INSERT INTO authors (name, email, description, created_at) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: FindAuthorByEmail :one
SELECT * FROM authors WHERE email = $1;