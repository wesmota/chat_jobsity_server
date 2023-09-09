package chatrooms

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/models"
	stmodels "github.com/wesmota/go-jobsity-chat-server/storage/models"
)

func (r *Repo) ListChatRooms(ctx context.Context) ([]models.ChatRoom, error) {
	var chatRooms []stmodels.ChatRoom
	var modelChatRooms []models.ChatRoom
	err := r.DB().Debug().WithContext(ctx).Order("id DESC").Find(&chatRooms).Error
	if err != nil {
		log.Info().Msgf("Error listing chat rooms: %v", err)
		return nil, err
	}
	// transform to models
	for i := range chatRooms {
		modelChatRooms = append(modelChatRooms, models.ChatRoom{
			ID:   chatRooms[i].ID,
			Name: chatRooms[i].Name,
		})
	}
	log.Info().Interface("chatRooms", chatRooms).Interface("modelChatRooms", modelChatRooms).Msg("ListChatRooms")
	return modelChatRooms, nil
}

func (r *Repo) CreateChatRoom(ctx context.Context, chatRoom models.ChatRoom) error {
	// form models to stmodels
	stChatRoom := stmodels.ChatRoom{
		Name: chatRoom.Name,
	}
	return r.DB().Debug().WithContext(ctx).Create(&stChatRoom).Error
}
