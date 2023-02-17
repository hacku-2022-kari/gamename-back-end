package useDB

import (
	"log"
)

type Player struct {
	PlayerName string
	Icon       int
	Role       int
	Theme      string
	Hint       string
	IsDelete   bool
	Answer     string
}

type RoomPlayer struct {
	Roomid   string
	Playerid string
}

func AddPlayer(roomId string, playerName string, playerIcon int) string {

	player := Player{
		PlayerName: playerName,
		Icon:       playerIcon,
		Role:       0,
		Theme:      "notheme",
		Hint:       "nohint",
		IsDelete:   false,
		Answer:     "noanswer",
	}

	ctx, client := connnectDB()

	docRef, _, err := client.Collection("Player").Add(ctx, player)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	roomPlayer := RoomPlayer{
		Roomid:   roomId,
		Playerid: docRef.ID,
	}
	ref := client.Collection("RoomPlayer").NewDoc()
	_, _err := ref.Set(ctx, roomPlayer)
	if _err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", _err)
	}
	defer client.Close()
	return docRef.ID
}

// $body = @{
//     roomId = "cbBipgOwuA8wxu5XAXFW"
//     playerName = "testman"
// 		playerIcon = 3
// } | ConvertTo-Json

// Invoke-RestMethod -Method POST -Uri http://localhost:1323/addPlayer -Body $body -ContentType "application/json"
