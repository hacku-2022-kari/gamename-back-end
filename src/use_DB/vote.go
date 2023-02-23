package useDB

import (
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
)

func Vote(inputPlayerId string, roomId string) bool {

	ctx, client, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	playerRef := client.Collection("Player").Doc(inputPlayerId)
	_, err = playerRef.Update(ctx, []firestore.Update{
		{Path: "Vote", Value: firestore.Increment(1)},
	})

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	roomDoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	bytes, _ := json.Marshal(roomDoc.Data()["PaticNum"])
	var particNumInt int64
	err = json.Unmarshal(bytes, &particNumInt)
	if err != nil {
		return false
	}

	var addStep bool = false

	var sumVote int = 0
	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			return false
		}
		bytes, _ := json.Marshal(playerDoc.Data()["Vote"])
		var voteInt int64
		err = json.Unmarshal(bytes, &voteInt)
		if err != nil {
			return false
		}
		sumVote += int(voteInt)
		if sumVote == int(particNumInt) {
			addStep = true
		}
	}

	if addStep == true {
		_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
			{Path: "Step", Value: 9},
		})
		if err != nil {
			return false
		}
	}

	defer client.Close()
	return true
}

// $body = @{
// 	roomId = "WgBySaSIBvs92OsDdd4i"
//     playerId = "W8fAxy12FB8fGF9vysxy"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/vote -Body $body -ContentType "application/json"
