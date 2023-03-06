package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func IsModeWolf(roomId string) bool {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	roomDoc, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Println("failed to connect to database: ", err)
	}

	defer client.Close()
	return roomDoc.Data()["IsModeWolf"].(bool)
}
