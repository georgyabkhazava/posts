package comment

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) SaveComment(ctx context.Context, userId int64, text string, twitId int64) (int64, error) {
	query := `insert into comments (user_id, text, twit_id) values ($1, $2, $3) returning id;`

	var id int64
	if err := s.db.QueryRowContext(ctx, query, userId, text, twitId).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "query context")
	}

	return id, nil
}
