package models

type User struct {
	ID       uint   `json:"id,omit"`
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
