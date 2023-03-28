package connectDB

import (
	"context"
	"gamename-back-end/pkg/utils"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func ConnectDB(param string) (context.Context, *firestore.Client, error) {
	credentials_file_path := ""
	if param == "GET_RANDOM_THEME" {
		credentials_file_path = "path/to/serviceAccount.json"
	} else {
		id := utils.DistributeDB(param)
		credentials_file_path = "path/to/db_" + id + ".json"
	}

	ctx := context.Background()
	sa := option.WithCredentialsFile(credentials_file_path)
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}
