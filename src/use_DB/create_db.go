package useDB

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Room struct {
	Password string `json:"pa"`
	PaticNum int
	Phase    int
	Step     string
}

func connnectDB() (context.Context, *firestore.Client) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	config := &firebase.Config{ProjectID: "gotest-bc4c6"}
	app, err := firebase.NewApp(ctx, config, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	return ctx, client
}

func CreateRoom(pass string, particNum int, phase int, step string) {
	room := Room{
		Password: pass,
		PaticNum: particNum,
		Phase:    phase,
		Step:     step,
	}

	ctx, client := connnectDB()

	ref := client.Collection("Room").NewDoc()
	_, err := ref.Set(ctx, room)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()

}
