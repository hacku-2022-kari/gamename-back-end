package useDB

import (
	"log"
)

func GetResult(roomId string) interface{}{
	ctx, client, err := connectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)	
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)	
	}
	result := iter.Data()["Result"]
	defer client.Close()
	return result

}
