package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"

	"cloud.google.com/go/firestore"
)

func UpdateAnswer(answer string, roomId string, playerId string) bool {

	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Answer": answer,
	}, firestore.MergeAll)
	if _err != nil {
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

// $body = @{
//     roomId = "idkAj1Km0ACPCkQybbPD"
//		playerId = ""
//     answer = "ピカチュウ"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/update-answer -Body $body -ContentType "application/json;charset=UTF-8"
