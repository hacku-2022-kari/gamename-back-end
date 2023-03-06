package useDB

import (
	"log"
)

func GetRole(playerId string) interface{} {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}
	iter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}
	role := iter.Data()["Role"]
	defer client.Close()

	return role
}
