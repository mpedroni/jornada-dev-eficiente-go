package domain

import "time"

type Author struct {
	ID          int
	Name        string
	Email       string
	Description string
	CreatedAt   time.Time
}

func NewAuthor(
	name string,
	email string,
	description string,
) *Author {
	return &Author{
		ID:          -1,
		Name:        name,
		Email:       email,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
