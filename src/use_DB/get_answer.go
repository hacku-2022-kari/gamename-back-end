package useDB

import (
	"log"
)

func GetAnswer(roomId string) string {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	answer := iter.Data()["Answer"].(string)
	defer client.Close()
	return answer

}
