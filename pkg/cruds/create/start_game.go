package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
)

func StartGame(roomId string) bool {

	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, err = roomRef.Update(ctx, []firestore.Update{
		{Path: "Phase", Value: firestore.Increment(1)},
		{Path: "IsExitWolf", Value: false},
		{Path: "PeaceVote", Value: 0},
	})

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()

	if err != nil {
		return false
	}

	var playerList []string

	for _, rpDoc := range rpDocs {
		playerId := rpDoc.Data()["PlayerId"].(string)
		playerList = append(playerList, playerId)
	}

	for i := range playerList {
		j := rand.Intn(i + 1)
		playerList[i], playerList[j] = playerList[j], playerList[i]
	}

	var roleCount int = 1
	rand.Seed(time.Now().UnixNano())

	for i := range playerList {
		playerDoc := client.Collection("Player").Doc(playerList[i])
		if err != nil {
			return false
		}
		_, err := playerDoc.Set(ctx, map[string]interface{}{
			"Role": roleCount,
			"Wolf": false,
		}, firestore.MergeAll)

		if err != nil {
			return false
		}
		if roleCount != 3 {
			roleCount += 1
		}
	}

	roomDoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if roomDoc.Data()["IsModeWolf"].(bool) == true {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(3) != 0 {
			_, err = roomRef.Update(ctx, []firestore.Update{
				{Path: "IsExitWolf", Value: true},
			})

			for i := range playerList {
				j := rand.Intn(i + 1)
				playerList[i], playerList[j] = playerList[j], playerList[i]
			}

			playerDoc := client.Collection("Player").Doc(playerList[0])
			if err != nil {
				return false
			}
			_, err = playerDoc.Set(ctx, map[string]interface{}{
				"Wolf": true,
			}, firestore.MergeAll)

			if err != nil {
				return false
			}
		}
	}

	_, err = roomRef.Set(ctx, map[string]interface{}{
		"Step": 1,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}
	defer client.Close()
	return true
}
