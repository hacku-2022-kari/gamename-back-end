package useDB

import (
	"log"
)

func JudgementAnswer(roomId string) bool {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	answer := iter.Data()["IsCorrect"].(bool)
	defer client.Close()
	return answer

}
