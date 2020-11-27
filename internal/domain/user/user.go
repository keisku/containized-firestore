package user

// User .
type User interface {
	ID() ID
	Account() Account
}

type user struct {
	id      ID
	account Account
}

func (u user) ID() ID {
	return u.id
}

func (u user) Account() Account {
	return u.account
}

// NewUser .
func NewUser(id ID, account Account) User {
	return user{id, account}
}

// Users .
type Users []User
