package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetStep(roomId string) interface{} {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}

	step := iter.Data()["Step"] //TODO型チェックをおこなう

	return step

}

//http://localhost:1323/step?roomid=cbBipgOwuA8wxu5XAXFW
