package useDB

import (
	"log"
)

func PlayerList(roomId string) [][]interface{} {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	var playerList [][]interface{}
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("error getting RoomPlayer documents: %v\n", err)
	}

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}
		playerName := playerDoc.Data()["PlayerName"].(string)
		playerIcon := int(playerDoc.Data()["Icon"].(int64))
		playerList = append(playerList, []interface{}{playerName, playerIcon})
	}

	return playerList

}

//cbBipgOwuA8wxu5XAXFW
