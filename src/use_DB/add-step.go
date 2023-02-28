package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

func AddStep(roomId string) bool {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err := roomRef.Set(ctx, map[string]interface{}{
		"Step": 8,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	defer client.Close()
	return true

}
