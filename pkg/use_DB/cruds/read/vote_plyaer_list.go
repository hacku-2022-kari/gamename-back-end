package readDB

import (
	"encoding/json"
	"fmt"
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

// TODO: 構造体の命名の検討
type VotePlayerInfo struct {
	PlayerId   string `json:"playerid"`
	NickName   string `json:"nickname"`
	ParticIcon int    `json:"particIcon"`
	Text       string `json:"text"`
}

func VotePlayerList(roomId string) []VotePlayerInfo {
	ctx, client, _err := connectDB.ConnectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("error getting RoomPlayer documents: %v\n", err)
	}

	var votePlayerList []VotePlayerInfo

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}

		var addPlayer VotePlayerInfo

		addPlayer.PlayerId = playerID
		addPlayer.NickName = playerDoc.Data()["PlayerName"].(string)
		addPlayer.ParticIcon = int(playerDoc.Data()["Icon"].(int64))

		bytes, _ := json.Marshal(playerDoc.Data()["Role"])
		var roleInt int64
		err = json.Unmarshal(bytes, &roleInt)
		if err != nil {
			log.Fatalf("error getting Player document: %v\n", err)
		}

		if int(roleInt) == 1 {
			addPlayer.Text = playerDoc.Data()["Answer"].(string)
			votePlayerList = append([]VotePlayerInfo{addPlayer}, votePlayerList...)
		} else {
			addPlayer.Text = playerDoc.Data()["Hint"].(string)
			votePlayerList = append(votePlayerList, addPlayer)

		}

	}
	fmt.Println(votePlayerList)
	return votePlayerList

}
