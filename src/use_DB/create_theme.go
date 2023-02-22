package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

func CreateTheme(inputTheme string, playerId string, roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	docRef := client.Collection("Player").Doc(playerId)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"Theme":  inputTheme,
		"Status": true,
	}, firestore.MergeAll)
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	var okCount int = 1
	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}
		if int(playerDoc.Data()["Role"].(int)) != 1 {
			if rpDoc.Data()["Hint"].(string) != "no-hint" {
				okCount += 1
			}
		}
	}

	rdoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		return false
	}
	if okCount == rdoc.Data()["PaticNum"] {
		_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
			{Path: "Step", Value: 3},
		})
		if err != nil {
			return false
		}
	}
	defer client.Close()
	return true
}

// $body = @{

//     playerId = "RB6srVwHGZ8Jih3QwKZ5"

//     Theme= "pokemon"

// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createTheme -Body $body -ContentType "application/json"
