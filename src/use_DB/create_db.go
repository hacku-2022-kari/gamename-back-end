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
	PaticNum      int
	Theme         string
	Phase         int
	Step          int
	IsModeWolf    bool
	IsExitWolf    bool
	PeaceVote     int
	IsCorrectWolf bool
	Result        int
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

func CreateRoom(particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int, isCorrectWolf bool) string {

	room := Room{
		PaticNum:      particNum,
		Theme:         theme,
		Phase:         phase,
		Step:          step,
		IsModeWolf:    wolfMode,
		IsExitWolf:    isExitWolf,
		PeaceVote:     peaceVote,
		IsCorrectWolf: isCorrectWolf,
		Result:        1,
	}

	ctx, client := connnectDB()

	ref := client.Collection("Room").NewDoc()
	_, err := ref.Set(ctx, room)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
	return ref.ID
}

// $body = @{
// 	wolfMode = $True
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/create-room -Body $body -ContentType "application/json"
