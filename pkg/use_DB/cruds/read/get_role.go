package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetRole(playerId string) interface{} {
	ctx, client, _err := connectDB.ConnectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	iter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}
	role := iter.Data()["Role"]
	defer client.Close()

	return role
}
