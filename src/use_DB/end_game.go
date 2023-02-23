package useDB

import (
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
)

func EndGame(roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "Answer", Value: "no-answer"},
		{Path: "Step", Value: 0},
		{Path: "Theme", Value: "no-theme"},
	})

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	if err != nil {
		return false
	}

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}
		roomRef := client.Collection("Player").Doc(playerID)
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "Answer", Value: "no-answer"},
			{Path: "Step", Value: 0},
			{Path: "Theme", Value: "no-theme"},
		})
		}
	}

	rand.Seed(time.Now().UnixNano())
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Step": 1,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}
	defer client.Close()
	return true
}
