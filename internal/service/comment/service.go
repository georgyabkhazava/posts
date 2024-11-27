package comment

import (
	"context"
	"github.com/georgyabkhazava/posts/internal/domain"
)

type TwitStorage interface {
	GetTwitById(ctx context.Context, id int64) (domain.Twit, error)
}

type CommentStorage interface {
	SaveComment(ctx context.Context, userId int64, text string, twitId int64) (int64, error)
}

type Service struct {
	twitStorage    TwitStorage
	commentStorage CommentStorage
}

func New(twitStorage TwitStorage, commentStorage CommentStorage) *Service {
	return &Service{
		twitStorage:    twitStorage,
		commentStorage: commentStorage,
	}
}
func (s *Service) CreateComment(ctx context.Context, userId int64, text string, twitId int64) (int64, error) {
	_, err := s.twitStorage.GetTwitById(ctx, twitId)
	if err != nil {
		return 0, err
	}
	return s.commentStorage.SaveComment(ctx, userId, text, twitId)
}
