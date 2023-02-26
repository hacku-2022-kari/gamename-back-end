package useDB

import (
	"log"
)

func GetStep(roomId string) interface{} {
	ctx, client, err := connectDB()
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

//http://localhost:1323/step?roomid=cbBipgOwuA8wxu5XAXFW
