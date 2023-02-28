package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

type Player struct {
	PlayerName string
	Icon       int
	Role       int
	Theme      string
	Hint       string
	IsDelete   bool
	Answer     string
	Wolf       bool
	Vote       int
	Point      int
}

type RoomPlayer struct {
	RoomId   string
	PlayerId string
}

func AddPlayer(roomId string, playerName string, playerIcon int) string {
	player := Player{
		PlayerName: playerName,
		Icon:       playerIcon,
		Role:       0,
		Theme:      "no-theme",
		Hint:       "no-hint",
		IsDelete:   false,
		Answer:     "no-answer",
		Wolf:       false,
		Vote:       0,
		Point:      0,
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
