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
	roomDoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Answer":           "no-answer",
		"Step":             0,
		"HowToDecideTheme": 0,
		"IsCorrect":        false,
		"Theme":            "no-theme",
		"IsMoodWolf" :		 roomDoc.Data()["WolfMode"].(bool),
		"IsExitWolf" : 		false,
		"PeaceVote"  :		0,
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
		playerDoc,err := client.Collection("Player").Doc(playerID).Get(ctx)
		_, err = playerRef.Set(ctx, map[string]interface{}{
			"Answer":   "no-answer",
			"Hint":     "no-hint",
			"Role":     0,
			"IsDelete": false,
			"Theme":    "no-theme",
			"Point":	playerDoc.Data()["Point"].(int),
			"Wolf":		false,
			"Vote" : 	0,
		}, firestore.MergeAll)
		if err != nil {
			return false
		}
	}

	defer client.Close()
	return true
}

// $body = @{
//     roomId = "WgBySaSIBvs92OsDdd4i"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/initialize -Body $body -ContentType "application/json"
