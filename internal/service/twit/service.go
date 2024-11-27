package twit

import (
	"context"
	"github.com/georgyabkhazava/posts/internal/domain"
)

type TwitsStorage interface {
	SaveTwit(ctx context.Context, title string, text string, userId int64) (int64, error)
	GetTwits(ctx context.Context, userId int64) ([]domain.Twit, error)
	DeleteTwits(ctx context.Context, id int64, userId int64) error
	GetTwitById(ctx context.Context, id int64) (domain.Twit, error)
}

type Service struct {
	storage TwitsStorage
}

func New(storage TwitsStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateTwit(ctx context.Context, title string, text string, userId int64) (int64, error) {
	return s.storage.SaveTwit(ctx, title, text, userId)
}

func (s *Service) GetTwitsByUserId(ctx context.Context, userId int64) ([]domain.Twit, error) {
	return s.storage.GetTwits(ctx, userId)
}

func (s *Service) DeleteTwitById(ctx context.Context, id int64, userId int64) error {
	return s.storage.DeleteTwits(ctx, id, userId)
}

func (s *Service) GetTwitById(ctx context.Context, id int64) (domain.Twit, error) {
	return s.storage.GetTwitById(ctx, id)
}
