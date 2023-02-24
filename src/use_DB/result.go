package useDB

import (
	"log"
)

func GetResult(roomId string) interface{}{
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}
	result := iter.Data()["Result"]
	defer client.Close()
	return result

}
