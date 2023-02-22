package useDB

func IsRoomExit(id string) bool {
	ctx, client := connnectDB()
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
