package useDB

import (
	"log"
)

// TODO: 構造体の命名の検討
type PlayerNNNIcon struct {
	NickName   string `json:"nickname"`
	ParticIcon int    `json:"particIcon"`
}

func PlayerList(roomId string) []PlayerNNNIcon {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
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
		// playerName := playerDoc.Data()["PlayerName"].(string)
		// playerIcon := int(playerDoc.Data()["Icon"].(int64))

		var addPlayer PlayerNNNIcon
		addPlayer.NickName = playerDoc.Data()["PlayerName"].(string)
		addPlayer.ParticIcon = int(playerDoc.Data()["Icon"].(int64))
		playerList = append(playerList, addPlayer)

		// playerList = append(playerList, []interface{}{playerName, playerIcon})
	}

	return playerList

}

//cbBipgOwuA8wxu5XAXFW
