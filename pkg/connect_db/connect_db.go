package connectDB

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func ConnectDB() (context.Context, *firestore.Client, error) { //TODO この関数とcreateDBにある関数で出力が違うため要検討
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	config := &firebase.Config{ProjectID: "gotest-bc4c6"}
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}
