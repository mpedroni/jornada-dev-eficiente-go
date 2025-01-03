package domain

import "context"

type AuthorRepository interface {
	Create(ctx context.Context, author *Author) error
	GetByEmail(ctx context.Context, email string) (*Author, error)
}
