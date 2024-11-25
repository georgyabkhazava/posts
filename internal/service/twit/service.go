package twit

import "context"

type AuthStorage interface {
	SaveTwit(ctx context.Context, title string, text string, userId int64) (int64, error)
}

type Service struct {
	storage AuthStorage
}

func New(storage AuthStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateTwit(ctx context.Context, title string, text string, userId int64) (int64, error) {
	return s.storage.SaveTwit(ctx, title, text, userId)
}
