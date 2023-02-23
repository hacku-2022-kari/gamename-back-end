package useDB

import (
	"log"
)

func GetStep(roomId string) interface{} {
	ctx, client, err := connectDB()
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
