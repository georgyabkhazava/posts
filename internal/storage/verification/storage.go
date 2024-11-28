package verification

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

func (s *Storage) SaveCode(ctx context.Context, userID int64, code string) error {
	query := `insert into codes (user_id, code) values ($1, $2);`

	_, err := s.db.ExecContext(ctx, query, userID, code)
	if err != nil {
		return errors.Wrap(err, "query context")
	}

	return nil
}

func (s *Storage) GetCodeByUserID(ctx context.Context, userID int64) (string, error) {
	query := `select code from codes where user_id = $1;`

	var code string
	if err := s.db.QueryRowContext(ctx, query, userID).Scan(&code); err != nil {
		return "", errors.Wrap(err, "query context")
	}

	return code, nil
}
