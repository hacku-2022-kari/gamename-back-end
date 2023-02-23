package useDB

import (
	"cloud.google.com/go/firestore"
)

func EndGame(roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Answer":           "no-answer",
		"Step":             0,
		"HowToDecideTheme": 0,
		"IsCorrect":        false,
		"Theme":            "no-theme",
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	if err != nil {
		return false
	}

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerRef := client.Collection("Player").Doc(playerID)
		_, err = playerRef.Set(ctx, map[string]interface{}{
			"Answer":   "no-answer",
			"Hint":     "no-hint",
			"Role":     0,
			"IsDelete": false,
			"Theme":    "no-theme",
		}, firestore.MergeAll)
		if err != nil {
			return false
		}
	}

	defer client.Close()
	return true
}
