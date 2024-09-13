package author

import (
	"context"
	"errors"
)

var (
	ErrAuthorNotFound     = errors.New("author not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type AuthorService interface {
	Create(ctx context.Context, name, email, description string) (*Author, error)
}

type service struct {
	repo AuthorRepository
}

func NewAuthorService(repo AuthorRepository) AuthorService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, name, email, description string) (*Author, error) {
	existing, err := s.repo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, ErrAuthorNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	a := NewAuthor(name, email, description)

	err = s.repo.Save(ctx, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
