package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"UserName,omitempty"`
	Email    string `json:"Email,omitempty"`
	Password string `json:"Password,omitempty"`
}
