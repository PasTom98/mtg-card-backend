package models

type User struct {
	Id        int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
