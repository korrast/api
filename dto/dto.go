package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type CreateTableRequest struct {
	Title string `json:"title" binding:"required,min=1,max=100"`
}

type CreateColumnRequest struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color"`
}

type GetTablesResponse struct {
	Id    string `json: "id"`
	Title string `json: "title"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
