package usecase

import "github.com/wesmota/go-jobsity-chat-server/models"

type ChatRoom interface {
	ListChatRooms() ([]models.ChatRoom, error)
}
