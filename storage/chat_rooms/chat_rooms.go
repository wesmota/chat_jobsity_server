package chatrooms

import (
	"context"
	"errors"

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

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (stmodels.User, error) {
	var user stmodels.User
	err := r.DB().Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *Repo) CreateUser(ctx context.Context, user models.User) error {
	// form models to stmodels
	stUser := stmodels.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	return r.DB().Debug().WithContext(ctx).Create(&stUser).Error
}

func (r *Repo) CreateChatMessage(ctx context.Context, message models.ChatMessage) error {
	// form models to stmodels
	user, err := r.GetUserByEmail(ctx, message.ChatUser)
	if err != nil || user.ID == 0 {
		return errors.New("user not found")
	}
	stMessage := stmodels.Chat{
		Message:    message.ChatMessage,
		UserId:     user.ID,
		ChatRoomId: message.ChatRoomId,
	}
	return r.DB().Debug().WithContext(ctx).Create(&stMessage).Error

}
