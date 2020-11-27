package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

// DeleteInput .
type DeleteInput struct {
	UserID string
}

func (u usecase) Delete(ctx context.Context, in *DeleteInput) error {
	userID, err := user.NewIDFromString(in.UserID)
	if err != nil {
		return err
	}
	return u.userRepository.DeleteAccount(ctx, *userID)
}
