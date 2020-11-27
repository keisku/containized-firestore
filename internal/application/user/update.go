package user

import "context"

// UpdateInput .
type UpdateInput struct {
	UserID    string
	AccountID string
	Mail      string
}

func (u usecase) Update(ctx context.Context, in *UpdateInput) error {
	return nil
}
