package useDB

import (
	"cloud.google.com/go/firestore"
)

func IsCorrect(roomId string, isCorrect bool) bool {
	ctx, client, _err := connectDB()
	if _err != nil {
		return false
	}
	defer client.Close()
	_, err := client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "IsCorrect", Value: isCorrect},
	})
	if err != nil {
		return false
	}

	return true

}

// $body = @{
//     roomId = "4ZNlgKuuDC7TdYl4xnih"
// 	isCorrect = $False
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/is-correct -Body $body -ContentType "application/json"
