package author

type CreateAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
