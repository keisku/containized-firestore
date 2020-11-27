package user

import "context"

// Repository .
type Repository interface {
	CreateAccount(context.Context, Account) (*ID, error)
	GetAccount(context.Context, ID) (*Account, error)
	Load(context.Context) (Users, error)
	UpdateAccount(context.Context, ID, Account) error
	DeleteAccount(context.Context, ID) error
}
