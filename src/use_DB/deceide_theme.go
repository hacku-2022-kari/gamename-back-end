package useDB

import (
	"cloud.google.com/go/firestore"
)

func DecideTheme(roomId string, howToDecideTheme int) bool {

	ctx, client, err := connectDB()

	if err != nil {
		return false
	}
	roomRef := client.Collection("Room").Doc(roomId)
	_, _err := roomRef.Set(ctx, map[string]interface{}{
		"HowToDecideTheme": howToDecideTheme,
	}, firestore.MergeAll)
	if _err != nil {
		return false
	}
	var step int = 2
	if howToDecideTheme == 1 {
		step = 3
	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: step},
	})
	if err != nil {
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
