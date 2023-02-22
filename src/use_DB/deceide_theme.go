package useDB

import (
	"cloud.google.com/go/firestore"
)

func DecideTheme(roomId string, howToDecideTheme int) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	docRef := client.Collection("Room").Doc(roomId)
	_, _err := docRef.Set(ctx, map[string]interface{}{
		"HowToDecideTheme": howToDecideTheme,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}
	defer client.Close()
	return true
}

// $body = @{
//     roomId = "zjH7Si3lo3vjtcqJSaE1"
//     howToDecideTheme = 1
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/how-to-decide-theme -Body $body -ContentType "application/json"
