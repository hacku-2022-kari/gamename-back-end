package useDB

import (
	"log"
	"strconv"
)

func GetStep(roomId string) int {
	ctx, client, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}

	data := iter.Data()
	step ,_:= strconv.Atoi(data["Step"].(string))

	return step

}
//http://localhost:1323/step?roomid=cbBipgOwuA8wxu5XAXFW
