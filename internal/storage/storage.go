package storage

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

func (s *Storage) SaveUser(ctx context.Context, name string, passwordHash string) (int64, error) {
	query := `insert into users (name, password) values ($1, $2) returning id;`

	var id int64
	if err := s.db.QueryRowContext(ctx, query, name, passwordHash).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "query context")
	}

	return id, nil
}
