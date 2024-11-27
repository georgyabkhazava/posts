package registration

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/georgyabkhazava/posts/internal/domain"
)

type RegistrationStorage interface {
	SaveUser(ctx context.Context, name string, passwordHash string, email string) (int64, error)
	GetUserID(ctx context.Context, name string, passwordHash string) (int64, error)
}

type VerificationService interface {
	SendVerificationCode(ctx context.Context, userID int64, email string) error
}

type Service struct {
	storage RegistrationStorage
}

func New(s RegistrationStorage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) RegistrationUser(ctx context.Context, name string, password string, email string) (int64, error) {
	hash := generateMD5(password)
	id, err := s.storage.SaveUser(ctx, name, hash, email)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func generateMD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func (s *Service) LoginUser(ctx context.Context, name string, password string) (string, error) {
	hash := generateMD5(password)
	userId, err := s.storage.GetUserID(ctx, name, hash)
	if err != nil {
		return "", err
	}
	token, err := domain.GenerateAccessToken(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}
