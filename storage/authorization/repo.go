package authorization

import (
	"database/sql"

	"github.com/wesmota/go-jobsity-chat-server/logger"
	"github.com/wesmota/go-jobsity-chat-server/storage"
)

type Repo struct {
	*storage.Repo
	logger logger.Logger
}

func NewAuthorizationsRepo(conn *sql.DB, logger logger.Logger) (*Repo, error) {
	r, err := storage.NewRepo(conn, logger)
	if err != nil {
		return nil, err
	}
	return &Repo{r, logger}, nil
}
