package useDB

import (
	"fmt"
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
		// 指定したフィールドの値を取得する
		value, ok := data["password"]
		if ok && value == password {
			fmt.Println("ドキュメントが存在し、指定したフィールドの値が一致します")
			return true
		}
	}

	fmt.Println("ドキュメントが存在しない、または指定したフィールドの値が一致しません")
	return false

}
