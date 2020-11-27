package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

func (r repository) CreateAccount(ctx context.Context, account user.Account) (*user.ID, error) {
	userID := user.GenerateID()
	_, err := r.accountDocument(*userID).Set(ctx, accountData{
		AccountID: account.ID().String(),
		Mail:      account.Mail().String(),
	})
	if err != nil {
		return nil, err
	}
	return userID, nil
}
