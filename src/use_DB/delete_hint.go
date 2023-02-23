package useDB

import (
	"cloud.google.com/go/firestore"
)

func DeleteHint(hintList []string, roomId string) bool {
	ctx, client, err := connectDB()
	if err != nil {
		return false
	}
	for i := 0; i < len(hintList); i++ {

		docRef := client.Collection("Player").Where("Hint", "==", hintList[i])
		docs, err := docRef.Documents(ctx).GetAll()
		if err != nil {
			return false
		}
		for _, doc := range docs {
			playerID := doc.Ref.ID
			docRef := client.Collection("Player").Doc(playerID)
			_, _err := docRef.Set(ctx, map[string]interface{}{
				"Hint": "no-hint",
			}, firestore.MergeAll)
			if _err != nil {
				return false
			}
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
