package category

import (
	"context"
	"errors"
)

var (
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryAlreadyExists = errors.New("category with the given name already exists")
)

type CategoryService interface {
	Create(ctx context.Context, name string) (*Category, error)
}

type categoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) Create(ctx context.Context, name string) (*Category, error) {
	existing, err := s.repo.FindByName(ctx, name)
	if err != nil && !errors.Is(err, ErrCategoryNotFound) {
		return nil, err
	}

	if existing != nil {
		return nil, ErrCategoryAlreadyExists
	}

	c, err := NewCategory(name)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(ctx, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
