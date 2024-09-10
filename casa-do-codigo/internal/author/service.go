package author

import (
	"context"
	"fmt"
)

type AuthorService interface {
	Create(ctx context.Context, name, email, description string) (Author, error)
}

type service struct {
}

func NewAuthorService() AuthorService {
	return &service{}
}

func (s *service) Create(ctx context.Context, name, email, description string) (Author, error) {
	fmt.Println(name, email, description)
	return Author{}, nil
}
