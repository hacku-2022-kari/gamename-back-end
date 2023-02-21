package useDB

import (
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateTheme(inputTheme string, id string) bool {

	ctx, client, err := connectDB()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	docRef := client.Collection("Player").Doc(id)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"Theme": inputTheme,
	}, firestore.MergeAll)
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}
	defer client.Close()
	return true
}

// $body = @{
//     playerId = "RB6srVwHGZ8Jih3QwKZ5"
//     Theme= "pokemon"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/createTheme -Body $body -ContentType "application/json"
