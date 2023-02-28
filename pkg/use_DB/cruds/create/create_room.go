package useDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

type Room struct {
	PaticNum      int
	Theme         string
	Phase         int
	Step          int
	IsModeWolf    bool
	IsExitWolf    bool
	PeaceVote     int
	IsCorrectWolf bool
	Result        int
}

func CreateRoom(particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int, isCorrectWolf bool) string {

	room := Room{
		PaticNum:      particNum,
		Theme:         theme,
		Phase:         phase,
		Step:          step,
		IsModeWolf:    wolfMode,
		IsExitWolf:    isExitWolf,
		PeaceVote:     peaceVote,
		IsCorrectWolf: isCorrectWolf,
		Result:        1,
	}

	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	ref := client.Collection("Room").NewDoc()
	_, err = ref.Set(ctx, room)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
	return ref.ID
}
