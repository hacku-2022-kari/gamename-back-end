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
	_, err = roomRef.Set(ctx, map[string]interface{}{
		"HowToDecideTheme": howToDecideTheme,
	}, firestore.MergeAll)
	if err != nil {
		return false
	}

	var step int = 2
	if howToDecideTheme == 1 {
		roomRef := client.Collection("Room").Doc(roomId)
		_, err := roomRef.Set(ctx, map[string]interface{}{
			"Theme": GetRandomTheme(),
		}, firestore.MergeAll)
		if err != nil {
			return false
		}
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
//     roomId = "gdCSnyP2pm3Gqf7UCIA4"
//     howToDecideTheme = 1
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/how-decide-theme -Body $body -ContentType "application/json"
