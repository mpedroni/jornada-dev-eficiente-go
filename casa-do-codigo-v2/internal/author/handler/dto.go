package handler

type CreateAuthorRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email" email:"email must be a valid email address"`
	Description string `json:"description" validate:"required,max=400"`
}
