package category

import "time"

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
