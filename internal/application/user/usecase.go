package user

import (
	"context"

	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

// UseCase .
type UseCase interface {
	Create(context.Context, *CreateInput) (*CreateOutput, error)
	Get(context.Context, *GetInput) (*GetOutput, error)
	Load(context.Context, *LoadInput) (*LoadOutput, error)
	Update(context.Context, *UpdateInput) error
	Delete(context.Context, *DeleteInput) error
}

type usecase struct {
	userRepository user.Repository
}

// NewUseCase .
func NewUseCase(
	userRepository user.Repository,
) UseCase {
	return usecase{
		userRepository: userRepository,
	}
}
