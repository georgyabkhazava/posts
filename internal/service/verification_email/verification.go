package verification_email

import "context"

type VerificationStorage interface {
	SaveCode(ctx context.Context, userID int64, email string, code string) (int64, error)
}

type ServiceEmail struct {
	varificationstorage VerificationStorage
}

func (v *ServiceEmail) SendVerificationCode(ctx context.Context, userID int64, email string) error {

}
