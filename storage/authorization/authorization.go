package authorization

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
	stmodels "github.com/wesmota/go-jobsity-chat-server/storage/models"
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user stmodels.User
	err := r.DB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	userModel := models.User{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	return userModel, err
}

func (r *Repo) CreateUser(ctx context.Context, user models.User) error {
	stUserModel := stmodels.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	return r.DB().Debug().WithContext(ctx).Create(&stUserModel).Error
}
