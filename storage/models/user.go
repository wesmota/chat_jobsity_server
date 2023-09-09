package stmodels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	Email    string
	Password string
}
