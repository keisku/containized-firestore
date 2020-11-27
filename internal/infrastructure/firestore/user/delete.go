package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

func (r repository) DeleteAccount(ctx context.Context, id user.ID) error {
	_, err := r.accountDocument(id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
