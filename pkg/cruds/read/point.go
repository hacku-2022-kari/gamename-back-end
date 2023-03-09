package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"

	"cloud.google.com/go/firestore"
)

func PointCal(roomId string) bool {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		return false
	}
	defer client.Close()
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		return false
	}

	roomDocs, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		return false
	}

	var villagePoint int = 1
	var wolfPoint int = 1
	if roomDocs.Data()["IsCorrect"].(bool) == true {
		if roomDocs.Data()["IsCorrectWolf"].(bool) == true {
			villagePoint = 3
			wolfPoint = 0
		} else {
			villagePoint = 1
			wolfPoint = 1
		}
	} else {
		if roomDocs.Data()["IsCorrectWolf"].(bool) == true {
			villagePoint = 2
			wolfPoint = 1
		} else {
			villagePoint = 0
			wolfPoint = 5
		}
	}

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			return false
		}
		playerRef := client.Collection("Player").Doc(playerID)
		if playerDoc.Data()["Wolf"].(bool) == true {
			_, err = playerRef.Update(ctx, []firestore.Update{
				{Path: "Point", Value: firestore.Increment(wolfPoint)},
			})
		} else {
			_, err = playerRef.Update(ctx, []firestore.Update{
				{Path: "Point", Value: firestore.Increment(villagePoint)},
			})
		}
	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 11},
	})
	if err != nil {
		log.Println("failed to connect to database: ", err)
	}

	return true

}
