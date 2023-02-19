package useDB

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Player struct {
	PlayerName string
	Icon       int
	Role       int
	Theme      string
	Hint       string
	IsDelete   bool
	Answer     string
}

type RoomPlayer struct {
	Roomid   string
	Playerid string
}

func connectDB() (context.Context, *firestore.Client, error) {
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

func AddPlayer(roomId string, playerName string, playerIcon int) string {
	fmt.Println("OK")
	player := Player{
		PlayerName: playerName,
		Icon:       playerIcon,
		Role:       0,
		Theme:      "notheme",
		Hint:       "nohint",
		IsDelete:   false,
		Answer:     "noanswer",
	}

	ctx, client, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	docRef, _, err := client.Collection("Player").Add(ctx, player)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	fmt.Println("OK")
	roomPlayer := RoomPlayer{
		Roomid:   roomId,
		Playerid: docRef.ID,
	}
	fmt.Println("OK")
	ref := client.Collection("RoomPlayer").NewDoc()
	_, _err := ref.Set(ctx, roomPlayer)
	if _err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", _err)
	}
	fmt.Println("OK")
	defer client.Close()
	return docRef.ID
}

$body = @{
    roomId = "cbBipgOwuA8wxu5XAXFW"
    playerName = "testman"
		playerIcon = 3
} | ConvertTo-Json

Invoke-RestMethod -Method POST -Uri http://localhost:1323/addPlayer -Body $body -ContentType "application/json"
