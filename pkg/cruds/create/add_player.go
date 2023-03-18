package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	types "gamename-back-end/pkg/types"
	"log"

	"cloud.google.com/go/firestore"
)

func AddPlayer(roomId string, playerName string, playerIcon int) string {
	player := types.CreatePlayer{
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

	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	docRef, _, err := client.Collection("Player").Add(ctx, player)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "ParticNum", Value: firestore.Increment(1)},
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomPlayer := types.RoomPlayer{
		RoomId:   roomId,
		PlayerId: docRef.ID,
	}
	ref := client.Collection("RoomPlayer").NewDoc()
	_, err = ref.Set(ctx, roomPlayer)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
	return docRef.ID
}
