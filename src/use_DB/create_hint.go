package useDB

import (
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateHint(inputHint string, playerId string, roomId string) bool {

	ctx, client, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	docRef := client.Collection("Player").Doc(playerId)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"Hint": inputHint,
	}, firestore.MergeAll)
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	var addStep bool = true

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			return false
		}
		if playerDoc.Data()["Role"].(string) != "1" {
			if playerDoc.Data()["Hint"].(string) == "no-hint" {
				fmt.Println(playerID)
				addStep = false
			}
		}

	}

	rdoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		return false
	}
	fmt.Println("OK")
	fmt.Println(rdoc.Data()["PaticNum"])
	if addStep == true {
		fmt.Println("OK")
		_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
			{Path: "Step", Value: 4},
		})
		if err != nil {
			return false
		}
	}

	defer client.Close()
	return true
}

// $body = @{
// 	playerId = "FMJDJf3S6uGhEwfsf5qR"
// 	Hint= "黄色"
// 	roomId = "idkAj1Km0ACPCkQybbPD"
// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/create-hint -Body $body -ContentType "application/json;charset=UTF-8"
