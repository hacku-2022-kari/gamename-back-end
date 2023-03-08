package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"

	"cloud.google.com/go/firestore"
)

func DeleteHint(hintList []string, roomId string) bool {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		return false
	}
	for i := 0; i < len(hintList); i++ {

		docRef := client.Collection("Player").Doc(hintList[i])
		_, err = docRef.Set(ctx, map[string]interface{}{
			"IsDelete": true,
		}, firestore.MergeAll)
		if err != nil {
			return false
		}

	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 5},
	})
	if err != nil {
		return false
	}

	defer client.Close()
	return true
}
