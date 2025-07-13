package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type CreateTableRequest struct {
	Title string `json:"title" binding:"required,min=1,max=100"`
}

type CreateColumnRequest struct {
	TableID string `json:"table_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Color   string `json:"color"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
