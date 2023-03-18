package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

func ConnectDBForTest() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "doutankyohi-db-for-test"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}

func InitializeDatabase(ctx context.Context, client *firestore.Client) {
	collectionName := "Room"
	if _, err := client.Collection(collectionName).Doc("dummy").Create(ctx, map[string]interface{}{}); err != nil {
		log.Printf("Failed to create collection %q: %v", collectionName, err)
	}
	fmt.Printf("Successfully created collection %q\n", collectionName)
}

// https://firebase.google.com/docs/firestore/manage-data/delete-data?hl=ja#collections を参照
func DeleteCollection(ctx context.Context, client *firestore.Client, batchSize int, collectionName string) error {
	ref := client.Collection(collectionName)
	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
