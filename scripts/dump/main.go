package main

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/kskumgk63/containized-firestore/pkg/env"
)

func main() {
	ctx := context.Background()
	projectID, err := env.ProjectID()
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, collection := range []string{
		"accounts",
		// Add a collection name if it grows
	} {
		dumpCollections(ctx, client, collection)
	}
}

func dumpCollections(ctx context.Context, client *firestore.Client, collection string) {
	docs, err := client.Collection(collection).Documents(ctx).GetAll()
	if err != nil {
		fmt.Printf("failed to dump %s: %+v\n", collection, err)
		return
	}
	decoration := strings.Repeat("-", 35)
	fmt.Printf("%s\n    %s\n%s\n", decoration, collection, decoration)
	for i, doc := range docs {
		var data map[string]interface{}
		if err := doc.DataTo(&data); err != nil {
			fmt.Printf("failed to dump %s: %+v\n", collection, err)
			return
		}
		fmt.Printf("[%d] %+v: %+v\n", i, doc.Ref.ID, data)
	}
	fmt.Println("") // Add an empty line
}
