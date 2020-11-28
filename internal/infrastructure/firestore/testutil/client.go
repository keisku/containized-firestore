package testutil

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// WithFirestoreClient .
func WithFirestoreClient(ctx context.Context, t *testing.T, fn func(ctx context.Context, client *firestore.Client)) {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	projectID := string(b)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		t.Error(err.Error())
		return
	}
	collections, err := client.Collections(ctx).GetAll()
	if err != nil {
		t.Error(err.Error())
		return
	}
	deleteDocs(
		ctx,
		client,
		getDocs(ctx, collections),
	)

	fn(ctx, client)
}

func getDocs(ctx context.Context, cols []*firestore.CollectionRef) []*firestore.DocumentSnapshot {
	if len(cols) == 0 {
		return nil
	}
	result := []*firestore.DocumentSnapshot{}
	for _, c := range cols {
		docs, err := c.Documents(ctx).GetAll()
		if err != nil {
			panic(err)
		}
		result = append(result, docs...)
		for _, doc := range docs {
			doc.Ref.Delete(ctx)
			subCols, err := doc.Ref.Collections(ctx).GetAll()
			if err != nil {
				panic(err)
			}
			docs := getDocs(ctx, subCols)
			result = append(result, docs...)
		}
	}
	return result
}

func deleteDocs(ctx context.Context, client *firestore.Client, docs []*firestore.DocumentSnapshot) {
	if len(docs) == 0 {
		return
	}
	const deleteLimit = 500
	for i := deleteLimit; len(docs) > 0; {
		if len(docs) < deleteLimit {
			i = len(docs)
		}
		batch := client.Batch()
		for _, doc := range docs[:i] {
			batch.Delete(doc.Ref)
		}
		_, err := batch.Commit(ctx)
		if err != nil {
			panic(err)
		}
		docs = docs[i:]
	}
}
