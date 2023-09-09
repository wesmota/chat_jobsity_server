package usecase

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/usecase/storage"
)

type ChatRoomService struct {
	repo storage.ChatRoomRepo
}

func NewChatService(repo storage.ChatRoomRepo) *ChatRoomService {
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

func (s *ChatRoomService) CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) error {
	return s.repo.CreateChatMessage(ctx, chatMessage)
}

func (s *ChatRoomService) CreateUser(ctx context.Context, user models.User) error {
	return s.repo.CreateUser(ctx, user)
}
