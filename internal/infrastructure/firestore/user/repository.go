package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kskumgk63/containized-firestore/internal/domain/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repository struct {
	*firestore.Client
}

// NewRepository .
func NewRepository(firestoreClient *firestore.Client) user.Repository {
	return &repository{firestoreClient}
}

func (r repository) accountCollection() *firestore.CollectionRef {
	return r.Collection("account")
}

func (r repository) accountDocument(id user.ID) *firestore.DocumentRef {
	return r.accountCollection().Doc(id.String())
}

func (r repository) accountSnapshots(ctx context.Context) ([]*firestore.DocumentSnapshot, error) {
	snaps, err := r.accountCollection().Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return snaps, nil
}

func (r repository) accountDocumentsByIDs(ids user.IDs) []*firestore.DocumentRef {
	docs := make([]*firestore.DocumentRef, len(ids))
	for i, id := range ids {
		docs[i] = r.accountDocument(id)
	}
	return docs
}

func (r repository) accountSnapshotsByIds(
	ctx context.Context,
	ids user.IDs,
) ([]*firestore.DocumentSnapshot, error) {
	snaps, err := r.GetAll(ctx, r.accountDocumentsByIDs(ids))
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}
	return snaps, nil
}
