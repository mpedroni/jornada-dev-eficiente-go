-- name: CreateBook :one
INSERT INTO books 
(
  title,
  abstract,
  table_of_content,
  price,
  number_of_pages,
  isbn,
  publish_date,
  category_id,
  author_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: FindBookByTitle :one
SELECT 
  b.*,
  sqlc.embed(c)
FROM 
  books b
JOIN 
  categories c ON b.category_id = c.id
WHERE 
  title = $1;

-- name: ListBooks :many
SELECT 
  b.*,
  sqlc.embed(c)
FROM 
  books b
JOIN 
  categories c ON b.category_id = c.id
ORDER BY 
  b.created_at DESC;