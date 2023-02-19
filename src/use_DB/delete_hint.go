package useDB

import (
	"log"
)

func DeleteHint(hintList []string) bool {

	ctx, client, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	for i := 0; i < len(hintList); i++ {
		docRef := client.Collection("Player").Where("Hint", "==", hintList[i])
		docs, err := docRef.Documents(ctx).GetAll()
		if err != nil {
			log.Fatalf("error getting RoomPlayer documents: %v\n", err)
		}
		for _, doc := range docs {
			playerID := doc.Data()["Playerid"].(string)
			CreateHint("NoHint", playerID)
		}
	}
	defer client.Close()
	return true
}

// $body = @{
//     playerId = "RB6srVwHGZ8Jih3QwKZ5"
//     Hint = '"黄色"'
// } | ConvertTo-Json -Depth 100

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createHint -Body $body -ContentType "application/json;charset=UTF-8"
