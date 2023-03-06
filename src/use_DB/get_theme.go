package useDB

import (
	"log"
)

func GetTheme(roomId string) string {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	theme := iter.Data()["Theme"].(string)
	defer client.Close()
	return theme

}
