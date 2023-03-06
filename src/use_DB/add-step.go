package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

func AddStep(roomId string) bool {
	ctx, client, err := connectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, nil := roomRef.Set(ctx, map[string]interface{}{
		"Step": 8,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	defer client.Close()
	return true

}
