package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

// TODO: 構造体の命名の検討
type PlayerNNNIcon struct {
	NickName   string `json:"nickname"`
	ParticIcon int    `json:"particIcon"`
}

func PlayerList(roomId string) []PlayerNNNIcon {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	var playerList []PlayerNNNIcon

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		var addPlayer PlayerNNNIcon
		addPlayer.NickName = playerDoc.Data()["PlayerName"].(string)
		addPlayer.ParticIcon = int(playerDoc.Data()["Icon"].(int64))
		playerList = append(playerList, addPlayer)

	}

	return playerList

}
