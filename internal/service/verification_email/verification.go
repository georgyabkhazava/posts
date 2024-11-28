package verification_email

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

const password = "kjtuhzjsgrkhqlcm"

type VerificationStorage interface {
	SaveCode(ctx context.Context, userID int64, code string) error
}

type RegistrationStorage interface {
	GetEmailByUserID(ctx context.Context, userID int64) (string, error)
}

type ServiceEmail struct {
	verificationstorage VerificationStorage
	registrationstorage RegistrationStorage
}

func New(v VerificationStorage, r RegistrationStorage) *ServiceEmail {
	return &ServiceEmail{
		verificationstorage: v,
		registrationstorage: r,
	}
}

// получаем мыло по юзерайди
// генерируем код
// сохранить в базу
// отправить на почту

func (s *ServiceEmail) SendVerificationCode(ctx context.Context, userID int64) error {
	email, err := s.registrationstorage.GetEmailByUserID(ctx, userID)
	if err != nil {
		return err
	}
	var code string
	code = generateVerificationCode()

	err = s.verificationstorage.SaveCode(ctx, userID, code)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("Код подтверждения", "gikoskokos@yandex.ru", password, "smtp.yandex.ru")
	err = smtp.SendMail("smtp.yandex.ru:25", auth, "gikoskokos@yandex.ru", []string{email}, []byte(fmt.Sprintf("Код подтверждения: %s", code)))
	if err != nil {
		return err
	}
	return nil

}

func generateVerificationCode() string {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var number int
	number = r.Intn(10000)
	result := strconv.Itoa(number)
	for len(result) < 4 {
		result = "0" + result
	}
	return result
}
