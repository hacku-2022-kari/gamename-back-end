package useDB

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Player struct {
	PlayerName string
	Icon       int
	Role       string
	Theme      string
	Hint       string
	IsDelete   bool
	Answer     string
}

type RoomPlayer struct {
	RoomId   string
	PlayerId string
}

func connectDB() (context.Context, *firestore.Client, error) { //TODO この関数とcreateDBにある関数で出力が違うため要検討
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
	player := Player{
		PlayerName: playerName,
		Icon:       playerIcon,
		Role:       "0",
		Theme:      "no-theme",
		Hint:       "no-hint",
		IsDelete:   false,
		Answer:     "no-answer",
	}

	ctx, client, err := connectDB()

	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	docRef, _, err := client.Collection("Player").Add(ctx, player)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "PaticNum", Value: firestore.Increment(1)},
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomPlayer := RoomPlayer{
		RoomId:   roomId,
		PlayerId: docRef.ID,
	}
	ref := client.Collection("RoomPlayer").NewDoc()
	_, _err := ref.Set(ctx, roomPlayer)
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}
	defer client.Close()
	return docRef.ID
}

// $body = @{
//     roomId = "idkAj1Km0ACPCkQybbPD"
//     playerName = "まえだ"
// 	playerIcon = 3
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/add-player -Body $body -ContentType "application/json;charset=UTF-8"
