package useDB

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Room struct {
	PaticNum   int
	Theme      string
	Phase      int
	Step       int
	IsModeWolf bool
	IsExitWolf bool
	PeaceVote  int
}

// DB を分散するための乱数
var randInt int = rand.Intn(5)
var randString string = strconv.Itoa(randInt)

func connectDB() (context.Context, *firestore.Client, error) { //TODO この関数とcreateDBにある関数で出力が違うため要検討
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/key-" + randString + ".json")
	config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID-" + randString)}
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

func CreateRoom(particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int) string {

	room := Room{
		PaticNum:   particNum,
		Theme:      theme,
		Phase:      phase,
		Step:       step,
		IsModeWolf: wolfMode,
		IsExitWolf: isExitWolf,
		PeaceVote:  peaceVote,
	}

	ctx, client, _ := connectDB()

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
