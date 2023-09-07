package storage

import (
	"database/sql"
	log "log"

	"os"
	"time"

	l "github.com/wesmota/go-jobsity-chat-server/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

// Repo ...
type Repo struct {
	db     *gorm.DB
	logger l.Logger
}

// NewRepo returns a repo
func NewRepo(conn *sql.DB, logger l.Logger) (*Repo, error) {
	// copied from the default here https://github.com/go-gorm/gorm/blob/master/logger/logger.go#L66
	gormLogger := gl.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gl.Config{
		SlowThreshold:             1000 * time.Millisecond,
		LogLevel:                  gl.Error,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
	})
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{Logger: gormLogger})

	if err != nil {
		return nil, err
	}

	return &Repo{db: db, logger: logger}, nil
}

// DB exposes the GORM database handle
func (repo *Repo) DB() *gorm.DB {
	return repo.db
}
