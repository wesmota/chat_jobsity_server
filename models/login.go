package models

type Login struct {
	User
	Token string `json:"token"`
}
