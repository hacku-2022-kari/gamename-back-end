package useDB

import (
	"cloud.google.com/go/firestore"
)

func UpdateAnswer(answer string, roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	docRef := client.Collection("Room").Doc(roomId)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"Answer": answer,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}
	defer client.Close()
	return true
}

// $body = @{
//     roomId = "4ZNlgKuuDC7TdYl4xnih"
//     answer = "ピカチュウ"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/update-answer -Body $body -ContentType "application/json;charset=UTF-8"
