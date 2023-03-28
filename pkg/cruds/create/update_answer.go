package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"

	"cloud.google.com/go/firestore"
)

func UpdateAnswer(answer string, roomId string, playerId string) bool {

	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Set(ctx, map[string]interface{}{
		"Answer": answer,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	playerRef := client.Collection("Player").Doc(playerId)
	_, err = playerRef.Set(ctx, map[string]interface{}{
		"Answer": answer,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 6},
	})
	if err != nil {
		return false
	}
	defer client.Close()
	return true
}
