package useDB

import "log"

func IsRoomExit(id string) bool {
	ctx, client, err := connectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	docRef := client.Collection("Room").Doc(id)

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		return false
	}

	if docSnapshot.Exists() { //TODO ここが必要かの検証
		return true
	}

	return false
}
