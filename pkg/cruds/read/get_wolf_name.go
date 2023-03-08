package readDB
import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func WolfName(roomId string) string {
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

	for _, rpDoc := range rpDocs {
		playerID := rpDoc.Data()["PlayerId"].(string)
		playerDoc, err := client.Collection("Player").Doc(playerID).Get(ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
		}

		if playerDoc.Data()["Wolf"].(bool) == true {
			return playerDoc.Data()["PlayerName"].(string)

		}
	}
	return "NoWolf"
}
