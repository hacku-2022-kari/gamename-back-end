package useDB

import (
	"log"
)

func IsRoomExit(id string, password string) bool {
	ctx, client := connnectDB()
	docRef := client.Collection("Room").Doc(id)

	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get document: %v", err)
	}

	if docSnapshot.Exists() {
		data := docSnapshot.Data()
		value, ok := data["password"]
		if ok && value == password {
			return true
		}
	}
	return false

}
