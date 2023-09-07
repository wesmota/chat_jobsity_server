package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Message    string
	UserId     uint     `gorm:"index"`
	User       User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	ChatRoomId uint     `gorm:"index"`
	ChatRoom   ChatRoom `gorm:"foreignKey:ChatRoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}
