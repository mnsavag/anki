package sqlite

import (
	"database/sql"

	"github.com/mnsavag/anki.git/internal/lib/log"
)

type SqliteRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewSqliteRepository(db *sql.DB, logger log.Logger) *SqliteRepository {
	return &SqliteRepository{
		db:     db,
		logger: logger,
	}
}
