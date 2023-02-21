package useDB

import (
	"log"

	"cloud.google.com/go/firestore"
)

func CreateHint(inputHint string, id string) bool {

	ctx, client, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	docRef := client.Collection("Player").Doc(id)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"Hint": inputHint,
	}, firestore.MergeAll)
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	return true
}

// $body = @{
//     playerId = "RB6srVwHGZ8Jih3QwKZ5"
//     Hint = '"黄色"'
// } | ConvertTo-Json -Depth 100

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createHint -Body $body -ContentType "application/json;charset=UTF-8"
