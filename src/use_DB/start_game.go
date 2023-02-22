package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

func StartGame(roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		log.Printf("failed to connect to database: %v", err)
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "Phase", Value: firestore.Increment(1)},
	})
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("error getting RoomPlayer documents: %v\n", err)
	}
	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc := client.Collection("Player").Doc(playerID)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}
		_, _err := playerDoc.Set(ctx, map[string]interface{}{
			"Role": 1,
		}, firestore.MergeAll)

		if _err != nil {
			log.Fatalf("failed to connect to database: %v", _err)
		}
	}
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"Step": 1,
	}, firestore.MergeAll)
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	return true
}

// $body = @{
//     roomId = "zjH7Si3lo3vjtcqJSaE1"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/start-game -Body $body -ContentType "application/json"
