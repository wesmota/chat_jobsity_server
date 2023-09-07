package authorization

import (
	"context"

	stmodels "github.com/wesmota/go-jobsity-chat-server/storage/models"
)

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (stmodels.User, error) {
	var user stmodels.User
	err := r.DB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *Repo) CreateUser(ctx context.Context, user stmodels.User) error {
	return r.DB().WithContext(ctx).Create(&user).Error
}
