package author

import (
	"context"
	"fmt"
)

type Service interface {
	CreateAuthor(ctx context.Context, name, email, description string) (Author, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) CreateAuthor(ctx context.Context, name, email, description string) (Author, error) {
	fmt.Println(name, email, description)
	return Author{}, nil
}
