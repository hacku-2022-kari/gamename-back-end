package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

// NOTE: roomId は ConnectDB の引数を増やすタイミングで使用するため一旦未使用で良い
func GetRole(playerId string, roomId string) interface{} {
	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	iter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}
	role := iter.Data()["Role"]
	defer client.Close()

	return role
}
