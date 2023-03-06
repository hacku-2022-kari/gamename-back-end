package readDB

import (
	"encoding/json"
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

type HintKey struct {
	PlayerId    string `json:"playerId"`
	AvatarIndex int    `json:"avatarIndex"`
	Hint        string `json:"hint"`
	IsDelete    bool   `json:"isDelete"`
}

func HintList(roomId string) []HintKey {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Println("error getting Player document: \n", err)
	}
	defer client.Close()
	var hintList []HintKey
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Println("error getting Player document: \n", err)
	}
	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
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
			var addHint HintKey
			addHint.PlayerId = playerID
			addHint.AvatarIndex = int(playerDoc.Data()["Icon"].(int64))
			addHint.Hint = playerDoc.Data()["Hint"].(string)
			addHint.IsDelete = bool(playerDoc.Data()["IsDelete"].(bool))
			hintList = append(hintList, addHint)
		}
	}

	return hintList

}

//cbBipgOwuA8wxu5XAXFW
