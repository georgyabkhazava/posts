package registration

import (
	"context"
	"database/sql"
	"github.com/georgyabkhazava/posts/internal/domain"
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

// GetUserID Возвращает id пользователя по имени и хэшу пароля
func (s *Storage) GetUserID(ctx context.Context, name string, passwordHash string) (int64, error) {
	query := `select id from users where name = $1 and password = $2 limit 1;`

	var userID int64
	row := s.db.QueryRowContext(ctx, query, name, passwordHash)
	err := row.Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, domain.ErrUserNotFound
		}

		return 0, err
	}

	return userID, nil
}
