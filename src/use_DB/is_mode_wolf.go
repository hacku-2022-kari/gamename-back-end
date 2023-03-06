package useDB

import (
	"log"
)

func IsModeWolf(roomId string) bool {
	ctx, client, err := connectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	roomDoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Println("failed to connect to database: ", err)
	}

	defer client.Close()
	return roomDoc.Data()["IsModeWolf"].(bool)
}
