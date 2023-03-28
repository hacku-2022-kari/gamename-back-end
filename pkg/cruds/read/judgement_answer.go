package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func JudgementAnswer(roomId string) bool {
	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	answer := iter.Data()["IsCorrect"].(bool)
	defer client.Close()
	return answer

}
