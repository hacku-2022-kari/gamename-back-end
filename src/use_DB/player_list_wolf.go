package useDB

import (
	"log"
)

// TODO: 構造体の命名の検討
type PlayerInfo struct {
	NickName   string `json:"nickname"`
	ParticIcon int    `json:"particIcon"`
	Point      int    `json:"point"`
}

func PlayerListWolf(roomId string) []PlayerInfo {
	ctx, client, err := connectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	var playerList []PlayerInfo

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		var addPlayer PlayerInfo
		addPlayer.NickName = playerDoc.Data()["PlayerName"].(string)
		addPlayer.ParticIcon = int(playerDoc.Data()["Icon"].(int64))
		addPlayer.Point = int(playerDoc.Data()["Point"].(int64))
		playerList = append(playerList, addPlayer)

	}

	return playerList

}
