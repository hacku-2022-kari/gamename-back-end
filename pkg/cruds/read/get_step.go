package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetStep(roomId string) interface{} {
	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	step := iter.Data()["Step"] //TODO型チェックをおこなう

	return step

}
