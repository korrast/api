package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type CreateTableRequest struct {
	Title string `json:"title" binding:"required,min=1,max=100"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (r *AuthRequest) Validate() error {
	return nil
}

func (r *CreateTableRequest) Validate() error {
	return nil
}
