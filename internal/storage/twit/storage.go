package twit

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

// SaveTwit Сохраняет твит в баззу данных
func (s *Storage) SaveTwit(ctx context.Context, title string, text string, userId int64) (int64, error) {
	query := `insert into twits (title, text, user_id) values ($1, $2, $3) returning id;`

	var id int64
	if err := s.db.QueryRowContext(ctx, query, title, text, userId).Scan(&id); err != nil {
		return 0, errors.Wrap(err, "query context")
	}

	return id, nil
}

func (s *Storage) GetTwits(ctx context.Context, userId int64) ([]domain.Twit, error) {
	query := `select id, title, text, user_id from twits where user_id = $1;`
	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, errors.Wrap(err, "select twits")
	}
	defer rows.Close()

	var result []domain.Twit
	for rows.Next() {
		var twit domain.Twit
		if err := rows.Scan(
			&twit.Id,
			&twit.Title,
			&twit.Text,
			&twit.UserId,
		); err != nil {
			return nil, errors.Wrap(err, "scan")
		}

		result = append(result, twit)
	}

	return result, nil
}

func (s *Storage) DeleteTwits(ctx context.Context, id int64, userId int64) error {
	query := `delete from twits where id = $1 and user_id = $2;`
	_, err := s.db.ExecContext(ctx, query, id, userId)
	if err != nil {
		return errors.Wrap(err, "query context")
	}

	return nil
}
func (s *Storage) GetTwitById(ctx context.Context, id int64) (domain.Twit, error) {
	query := `select id, title, text, user_id from twits where id = $1 limit 1;`
	var twit domain.Twit
	if err := s.db.QueryRowContext(ctx, query, id).Scan(
		&twit.Id,
		&twit.Title,
		&twit.Text,
		&twit.UserId,
	); err != nil {
		return domain.Twit{}, errors.Wrap(err, "query context")
	}

	return twit, nil
}
