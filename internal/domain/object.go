package domain

type Object struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
