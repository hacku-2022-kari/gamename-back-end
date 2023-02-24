package useDB

import (
	"encoding/json"
	"log"
	"math/rand"

	"cloud.google.com/go/firestore"
)

// TODO: 構造体の命名の検討

func ChoiceWolf(roomId string) string {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Println("failed to connect to database: ", _err)
	}

	rpQuery := client.Collection("RoomPlayer").Where("RoomId", "==", roomId)
	rpDocs, err := rpQuery.Documents(ctx).GetAll()
	if err != nil {
		log.Println("failed to connect to database: ", _err)
	}

	var maxVote int = 0
	var choicedWolf []string

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Println("failed to connect to database: ", _err)
		}

		bytes, _ := json.Marshal(playerDoc.Data()["Vote"])
		var voteInt int64
		err = json.Unmarshal(bytes, &voteInt)
		if err != nil {
			log.Println("failed to connect to database: ", _err)
		}
		if maxVote == int(voteInt) {
			choicedWolf = append(choicedWolf, playerID)
		}
		if maxVote < int(voteInt) {
			maxVote = int(voteInt)
			choicedWolf = nil
			choicedWolf = append(choicedWolf, playerID)
		}

	}

	if len(choicedWolf) != 1 {
		for i := range choicedWolf {
			j := rand.Intn(i + 1)
			choicedWolf[i], choicedWolf[j] = choicedWolf[j], choicedWolf[i]
		}

	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 10},
	})
	if err != nil {
		log.Println("failed to connect to database: ", _err)
	}
	defer client.Close()
	return choicedWolf[0]

}
