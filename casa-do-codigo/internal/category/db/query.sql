-- name: CreateCategory :one
INSERT INTO categories (name, created_at) VALUES ($1, $2) RETURNING id;

-- name: FindCategoryByName :one
SELECT * FROM categories WHERE name = $1;