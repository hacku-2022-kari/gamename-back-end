package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"

	"cloud.google.com/go/firestore"
)

func IsCorrect(roomId string, isCorrect bool) bool {
	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		return false
	}

	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "IsCorrect", Value: isCorrect},
	})
	if err != nil {
		return false
	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 7},
	})
	if err != nil {
		return false
	}

	defer client.Close()
	return true

}
