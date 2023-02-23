package useDB

import (
	"cloud.google.com/go/firestore"
)

func IsCorrect(roomId string, isCorrect bool) bool {
	ctx, client, _err := connectDB()
	if _err != nil {
		return false
	}

	roomRef := client.Collection("Room").Doc(roomId)
	_, err := roomRef.Update(ctx, []firestore.Update{
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

// $body = @{
//     roomId = "idkAj1Km0ACPCkQybbPD"
// 	isCorrect = $False
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/is-correct -Body $body -ContentType "application/json"
