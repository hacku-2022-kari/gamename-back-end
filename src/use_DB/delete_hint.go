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

		docRef := client.Collection("Player").Where("Hint", "==", hintList[i])
		docs, err := docRef.Documents(ctx).GetAll()
		if err != nil {
			return false
		}
		for _, doc := range docs {
			playerID := doc.Ref.ID
			docRef := client.Collection("Player").Doc(playerID)
			_, _err := docRef.Set(ctx, map[string]interface{}{
				"Hint": "no-hint",
			}, firestore.MergeAll)
			if _err != nil {
				return false
			}
		}
	}

	roomRef := client.Collection("Room").Doc(roomId)
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Step": 5,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}

	defer client.Close()
	return true
}

// $body = @{
//     hint = @("黄色")
// } | ConvertTo-Json -Depth 100
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/delete-hint -Body $body -ContentType "application/json;charset=UTF-8"
