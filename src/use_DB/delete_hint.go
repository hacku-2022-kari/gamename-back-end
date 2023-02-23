package useDB

import (
	"cloud.google.com/go/firestore"
)

func DeleteHint(hintList []string, roomId string) bool {
	ctx, client, err := connectDB()
	if err != nil {
		return false
	}
	for i := 0; i < len(hintList); i++ {

		docRef := client.Collection("Player").Doc(hintList[i])
		_, _err := docRef.Set(ctx, map[string]interface{}{
			"Hint":     "同担拒否",
			"IsDelete": true,
		}, firestore.MergeAll)
		if _err != nil {
			return false
		}

	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 5},
	})
	if err != nil {
		return false
	}

	defer client.Close()
	return true
}

// $body = @{
// 	roomId = "cvi4EfisvGjd5jUJu3PS"
//     hint = @("OfjME4tAGeheTuUHLIQu")
// } | ConvertTo-Json -Depth 100
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/delete-hint -Body $body -ContentType "application/json;charset=UTF-8"
