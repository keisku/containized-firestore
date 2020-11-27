package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kskumgk63/containized-firestore/internal/domain/user"
)

func (r repository) GetAccount(ctx context.Context, userID user.ID) (*user.Account, error) {
	snap, err := r.accountDocument(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	return r.convertSnapToAccount(snap)
}

func (r repository) Load(ctx context.Context) (user.Users, error) {
	snaps, err := r.accountSnapshots(ctx)
	if err != nil {
		return nil, err
	}
	return r.convertSnapsToUsers(snaps)
}

func (r repository) convertSnapsToUsers(snaps []*firestore.DocumentSnapshot) (user.Users, error) {
	users := make(user.Users, len(snaps))
	for i, snap := range snaps {
		account, err := r.convertSnapToAccount(snap)
		if err != nil {
			return nil, err
		}
		userID, err := user.NewIDFromString(snap.Ref.ID)
		if err != nil {
			return nil, err
		}
		users[i] = user.NewUser(*userID, *account)
	}
	return users, nil
}

func (r repository) convertSnapToAccount(snap *firestore.DocumentSnapshot) (*user.Account, error) {
	var data accountData
	if err := snap.DataTo(&data); err != nil {
		return nil, err
	}
	accountID, err := user.NewAccountID(data.AccountID)
	if err != nil {
		return nil, err
	}
	mail, err := user.NewMail(data.Mail)
	if err != nil {
		return nil, err
	}
	return user.NewAccount(*accountID, *mail), nil
}
