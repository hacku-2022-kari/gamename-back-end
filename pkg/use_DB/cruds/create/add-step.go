package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"

	"cloud.google.com/go/firestore"
)

func AddStep(roomId string) bool {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Set(ctx, map[string]interface{}{
		"Step": 8,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	defer client.Close()
	return true

}
