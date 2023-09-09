package stmodels

import "gorm.io/gorm"

type ChatRoom struct {
	gorm.Model
	Name string
}
