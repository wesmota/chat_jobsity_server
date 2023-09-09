package chatrooms

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/usecase/storage"
)

type ChatRoomService struct {
	repo storage.ChatRoomRepo
}

func New(repo storage.ChatRoomRepo) *ChatRoomService {
	return &ChatRoomService{
		repo: repo,
	}
}

func (s *ChatRoomService) ListChatRooms(ctx context.Context) ([]models.ChatRoom, error) {
	return s.repo.ListChatRooms(ctx)
}

func (s *ChatRoomService) CreateChatRoom(ctx context.Context, chatRoom models.ChatRoom) error {
	return s.repo.CreateChatRoom(ctx, chatRoom)
}
