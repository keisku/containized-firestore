package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

// UpdateInput .
type UpdateInput struct {
	UserID    string
	AccountID string
	Mail      string
}

func (u usecase) Update(ctx context.Context, in *UpdateInput) error {
	userID, err := user.NewIDFromString(in.UserID)
	if err != nil {
		return err
	}
	accountID, err := user.NewAccountID(in.AccountID)
	if err != nil {
		return err
	}
	mail, err := user.NewMail(in.Mail)
	if err != nil {
		return err
	}
	account := user.NewAccount(*accountID, *mail)
	return u.userRepository.UpdateAccount(ctx, *userID, *account)
}
