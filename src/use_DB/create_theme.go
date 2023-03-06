package useDB

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
)

func CreateTheme(inputTheme string, playerId string, roomId string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	docRef := client.Collection("Player").Doc(playerId)
	_, err = docRef.Set(ctx, map[string]interface{}{
		"Theme":  inputTheme,
		"Status": true,
	}, firestore.MergeAll)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
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

		bytes, _ := json.Marshal(playerDoc.Data()["Role"])
		var roleInt int64
		err = json.Unmarshal(bytes, &roleInt)
		if err != nil {
			return false
		}
		if int(roleInt) != 1 {
			if playerDoc.Data()["Theme"].(string) == "no-theme" {
				addStep = false
			}
		}
	}

	if addStep == true {
		var playerList []string
		rand.Seed(time.Now().UnixNano())
		for _, rpDoc := range rpDocs {
			playerId := rpDoc.Data()["PlayerId"].(string)
			playerDoc, err := client.Collection("Player").Doc(playerId).Get(ctx)
			if err != nil {
				log.Println("error getting Player document: \n", err)
			}
			bytes, _ := json.Marshal(playerDoc.Data()["Role"])
			var roleInt int64
			err = json.Unmarshal(bytes, &roleInt)
			if err != nil {
				log.Println("error getting Player document: \n", err)
			}
			if int(roleInt) != 1 {
				playerList = append(playerList, playerId)
			}
		}
		num := rand.Intn(len(playerList))
		playerDoc, err := client.Collection("Player").Doc(playerList[num]).Get(ctx)
		if err != nil {
			return false
		}

		roomRef := client.Collection("Room").Doc(roomId)
		_, _err := roomRef.Set(ctx, map[string]interface{}{
			"Theme": playerDoc.Data()["Theme"].(string),
		}, firestore.MergeAll)
		if _err != nil {
			return false
		}
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
