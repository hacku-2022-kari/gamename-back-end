package useDB

import (
	"fmt"
	"log"
)

func IsRoomExit(id string, password string) bool {
	ctx, client := connnectDB()
	docRef := client.Collection("your-collection").Doc(id)
	query := docRef.Collection("room").Where("password", "==", password)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Failed to get document: %v", err)
	}

	if len(docs) > 0 {
		fmt.Println("ドキュメントが存在します")
		defer client.Close()
		return true
	} else {

		fmt.Println("ドキュメントが存在しません")
		defer client.Close()
		return false
	}

}
