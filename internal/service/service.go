package service

import (
	"context"
	"crypto/md5"
	"fmt"
)

type RegistrationStorage interface {
	SaveUser(ctx context.Context, name string, passwordHash string) (int64, error)
}

type Service struct {
	storage RegistrationStorage
}

func New(s RegistrationStorage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) RegistrationUser(ctx context.Context, name string, password string) (int64, error) {
	hash := generateMD5(password)
	id, err := s.storage.SaveUser(ctx, name, hash)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func generateMD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
