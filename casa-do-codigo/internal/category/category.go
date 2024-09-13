package category

import (
	"context"
	"time"
)

type CategoryRepository interface {
	Save(ctx context.Context, c *Category) error
	FindByName(ctx context.Context, name string) (*Category, error)
}

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	return &Category{
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}
