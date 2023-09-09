package storage

import (
	"context"

	"github.com/wesmota/go-jobsity-chat-server/models"
)

type AuthorizationRepo interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
}
