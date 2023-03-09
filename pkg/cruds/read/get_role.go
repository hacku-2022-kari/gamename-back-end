package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetRole(playerId string) interface{} {
	ctx, client, err := connectDB.ConnectDB()
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
