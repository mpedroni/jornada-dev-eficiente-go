package author

type CreateAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description" binding:"required,max=400"`
}
