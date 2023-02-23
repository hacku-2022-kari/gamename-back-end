package useDB

import (
	"log"
)

type HintKey struct {
	AvatorIndex int    `json:"avatorIndex"`
	Hint        string `json:"hint"`
	IsDelete    bool   `json:"isDelete"`
}

func HintList(roomId string) []HintKey {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	var hintList []HintKey
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
		var addHint HintKey
		addHint.AvatorIndex = int(playerDoc.Data()["Icon"].(int64))
		addHint.Hint = playerDoc.Data()["Hint"].(string)
		addHint.IsDelete = bool(playerDoc.Data()["IsDelete"].(bool))
		hintList = append(hintList, addHint)
	}

	return hintList

}

//cbBipgOwuA8wxu5XAXFW
