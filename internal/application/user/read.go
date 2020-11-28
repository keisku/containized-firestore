package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

// GetInput .
type GetInput struct {
	UserID string
}

// GetOutput .
type GetOutput struct {
	UserID    string
	AccountID string
	Mail      string
}

func (u usecase) Get(ctx context.Context, in *GetInput) (*GetOutput, error) {
	userID, err := user.NewIDFromString(in.UserID)
	if err != nil {
		return nil, err
	}
	account, err := u.userRepository.GetAccount(ctx, *userID)
	if err != nil {
		return nil, err
	}
	return &GetOutput{
		UserID:    userID.String(),
		AccountID: account.ID().String(),
		Mail:      account.Mail().String(),
	}, nil
}

// LoadInput .
type LoadInput struct {
}

// LoadOutput .
type LoadOutput struct {
	Users []LoadedUser
}

// LoadedUser .
type LoadedUser struct {
	ID        string
	AccountID string
	Mail      string
}

func (u usecase) Load(ctx context.Context, in *LoadInput) (*LoadOutput, error) {
	loadedUsers, err := u.userRepository.Load(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]LoadedUser, len(loadedUsers))
	for i, loadedUser := range loadedUsers {
		users[i] = LoadedUser{
			ID:        loadedUser.ID().String(),
			AccountID: loadedUser.Account().ID().String(),
			Mail:      loadedUser.Account().Mail().String(),
		}
	}
	return &LoadOutput{users}, nil
}
