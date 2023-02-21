package useDB

import (
	"log"
)

func HintList(roomId string) [][]interface{} {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	var hintList [][]interface{}
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
		playerHint:= playerDoc.Data()["Hint"]
		playerIsDelete := playerDoc.Data()["isDelete"]
		hintList = append(hintList, []interface{}{playerID, playerHint,playerIsDelete})
	}

	return hintList

}

//cbBipgOwuA8wxu5XAXFW
