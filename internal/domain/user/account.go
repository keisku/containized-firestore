package user

import (
	"fmt"
	"regexp"

	"github.com/kskumgk63/containized-firestore/internal/domain/errors"
)

// Account .
type Account struct {
	id   AccountID
	mail Mail
}

// NewAccount .
func NewAccount(id AccountID, mail Mail) *Account {
	return &Account{
		id:   id,
		mail: mail,
	}
}

// ID .
func (a Account) ID() AccountID {
	return a.id
}

// Mail .
func (a Account) Mail() Mail {
	return a.mail
}

// Accounts .
type Accounts []Account

// AccountID .
type AccountID struct {
	v string
}

// NewAccountID .
func NewAccountID(v string) (*AccountID, error) {
	if v == "" {
		return nil, fmt.Errorf("account_id must be set: %+v", errors.BadParams)
	}
	return &AccountID{v}, nil
}

// String .
func (id AccountID) String() string {
	return id.v
}

// Mail .
type Mail struct {
	v string
}

// HTML5 Compliant
// https://www.w3.org/TR/html5/forms.html#valid-e-mail-address
var emailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// NewMail .
func NewMail(v string) (*Mail, error) {
	if emailRegExp.MatchString(v) {
		return &Mail{v}, nil
	}
	return nil, fmt.Errorf("%s does not match mail format: %+v", v, errors.BadParams)
}

// String .
func (m Mail) String() string {
	return m.v
}
