package useDB

import (
	"fmt"
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

	var addStep bool = true

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			fmt.Println("OK1")
			return false
		}

		if playerDoc.Data()["Role"].(string) != "1" { //TODO int(playerDoc.Data()["Role"](int))これでいけないのはなぜ
			if playerDoc.Data()["Theme"].(string) == "no-theme" {
				fmt.Println(playerDoc.Data()["Theme"].(string))
				fmt.Println("OK")
				addStep = false
			}
		}
	}

	if addStep == true {
		_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
			{Path: "Step", Value: 3},
		})
		if err != nil {
			fmt.Println("OK")
			return false
		}
	}

	defer client.Close()
	return true
}

func isEqualToInt(i interface{}, x int) bool {
	if v, ok := i.(int); ok {
		return v == x
	}
	return false
}

// $body = @{
// 	playerId = "FMJDJf3S6uGhEwfsf5qR"
// 	Theme= "ポケモン"
// 	roomId = "idkAj1Km0ACPCkQybbPD"
// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/create-theme -Body $body -ContentType "application/json;charset=UTF-8"
