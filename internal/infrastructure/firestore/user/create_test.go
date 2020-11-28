package user

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/kskumgk63/containized-firestore/internal/domain/user"
	"github.com/kskumgk63/containized-firestore/internal/infrastructure/firestore/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_repository_CreateAccount(t *testing.T) {
	var (
		accountID, _ = user.NewAccountID("test")
		mail, _      = user.NewMail("test@example.com")
	)
	type args struct {
		ctx     context.Context
		account user.Account
	}
	tests := []struct {
		name    string
		args    args
		want    *user.ID
		wantErr string
	}{
		{
			name: "success",
			args: args{
				ctx:     context.Background(),
				account: *user.NewAccount(*accountID, *mail),
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.WithFirestoreClient(tt.args.ctx, t, func(ctx context.Context, client *firestore.Client) {
				r := repository{client}
				got, err := r.CreateAccount(tt.args.ctx, tt.args.account)
				if err != nil {
					assert.EqualError(t, err, tt.wantErr)
					return
				}
				assert.NoError(t, err)
				snap, _ := r.accountDocument(*got).Get(ctx)
				var data accountData
				snap.DataTo(&data)
				if !snap.Exists() {
					t.Error("created document does not exist")
				}
			})
		})
	}
}
