package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetAnswer(roomId string) string {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Println("error getting Player document: \n", err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Println("error getting Player document: \n", err)
	}
	answer := iter.Data()["Answer"].(string)
	defer client.Close()
	return answer

}
