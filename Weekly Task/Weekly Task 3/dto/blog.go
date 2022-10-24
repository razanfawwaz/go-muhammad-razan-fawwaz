package dto

type BlogRequest struct {
	ID      int
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}
