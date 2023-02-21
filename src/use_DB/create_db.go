package useDB

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Room struct {
	Password string
	PaticNum int
	Theme    string
	Phase    int
	Step     int
}

func connnectDB() (context.Context, *firestore.Client) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccount.json")
	config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(ctx, config, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	return ctx, client
}

func CreateRoom(password string, particNum int, theme string, phase int, step int) {
	room := Room{
		Password: password,
		PaticNum: particNum,
		Theme:    theme,
		Phase:    phase,
		Step:     step,
	}

	ctx, client := connnectDB()

	ref := client.Collection("Room").NewDoc()
	_, err := ref.Set(ctx, room)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()

}
