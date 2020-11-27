package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

func (r repository) UpdateAccount(ctx context.Context, id user.ID, account user.Account) error {
	_, err := r.accountDocument(id).Set(ctx, accountData{
		AccountID: account.ID().String(),
		Mail:      account.Mail().String(),
	})
	if err != nil {
		return err
	}
	return nil
}
