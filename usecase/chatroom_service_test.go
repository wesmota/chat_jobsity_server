package usecase

import (
	"context"
	"testing"

	"github.com/wesmota/go-jobsity-chat-server/mocks"
	"github.com/wesmota/go-jobsity-chat-server/models"
)

func TestListRooms(t *testing.T) {
	cases := map[string]struct {
		mockRepo  func() *mocks.ChatRoomRepo
		expectErr error
	}{
		"success": {
			mockRepo: func() *mocks.ChatRoomRepo {
				mockRepo := mocks.NewChatRoomRepo(t)
				mockRepo.On("ListChatRooms", context.Background()).Return([]models.ChatRoom{}, nil)
				return mockRepo
			},
		},
		"error": {
			mockRepo: func() *mocks.ChatRoomRepo {
				mockRepo := mocks.NewChatRoomRepo(t)
				mockRepo.On("ListChatRooms", context.Background()).Return([]models.ChatRoom{}, nil)
				return mockRepo
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			repo := tt.mockRepo()
			service := NewChatService(repo)
			_, err := service.ListChatRooms(context.Background())
			if err != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}
		})
	}

}
