package storage

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
)

type ChatRoomRepo interface {
	CreateChatRoom(ctx context.Context, chatRoom models.ChatRoom) error
	ListChatRooms(ctx context.Context) ([]models.ChatRoom, error)
	CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) error
	CreateUser(ctx context.Context, user models.User) error
}
