package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

// CreateInput .
type CreateInput struct {
	AccountID string
	Mail      string
}

// CreateOutput .
type CreateOutput struct {
	UserID string
}

func (u usecase) Create(ctx context.Context, in *CreateInput) (*CreateOutput, error) {
	accountID, err := user.NewAccountID(in.AccountID)
	if err != nil {
		return nil, err
	}
	mail, err := user.NewMail(in.Mail)
	if err != nil {
		return nil, err
	}
	account := user.NewAccount(*accountID, *mail)
	id, err := u.userRepository.CreateAccount(ctx, *account)
	if err != nil {
		return nil, err
	}
	return &CreateOutput{
		UserID: id.String(),
	}, nil
}
