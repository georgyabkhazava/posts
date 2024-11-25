package twit

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

// SaveTwit Сохраняет твит в баззу данных
func (s *Storage) SaveTwit(ctx context.Context, title string, text string, userId int64) (int64, error) {
	query := `insert into twits (title, text, user_id) values ($1, $2, $3) returning id;`

	var id int64
	if err := s.db.QueryRowContext(ctx, query, title, text, userId).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "query context")
	}

	return id, nil
}
