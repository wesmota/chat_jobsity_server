package chatrooms

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
	stmodels "github.com/wesmota/go-jobsity-chat-server/storage/models"
)

func (r *Repo) ListChatRooms(ctx context.Context) ([]models.ChatRoom, error) {
	var chatRooms []stmodels.ChatRoom
	var modelChatRooms []models.ChatRoom
	err := r.DB().WithContext(ctx).Order("id DESC").Find(&chatRooms).Error
	// transform to models
	for i := range chatRooms {
		modelChatRooms = append(modelChatRooms, models.ChatRoom{
			ID:   chatRooms[i].ID,
			Name: chatRooms[i].Name,
		})
	}
	return modelChatRooms, err
}

func (r *Repo) CreateChatRoom(ctx context.Context, chatRoom models.ChatRoom) error {
	// form models to stmodels
	stChatRoom := stmodels.ChatRoom{
		Name: chatRoom.Name,
	}
	return r.DB().WithContext(ctx).Create(&stChatRoom).Error
}
