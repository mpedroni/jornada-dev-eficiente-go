// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package melisqlc

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (login, password, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id
`

type CreateUserParams struct {
	Login     string
	Password  string
	CreatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Login, arg.Password, arg.CreatedAt)
	var id int32
	err := row.Scan(&id)
	return id, err
}
