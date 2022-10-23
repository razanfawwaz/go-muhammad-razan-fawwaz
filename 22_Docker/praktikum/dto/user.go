package dto

type CreateUserRequest struct {
	ID       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserRequest struct {
	ID    int
	Name  string `json:"name"`
	Email string `json:"email"`
}
