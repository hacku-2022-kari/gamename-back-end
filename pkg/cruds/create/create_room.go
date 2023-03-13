package createDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	structManege "gamename-back-end/pkg/struct"
	"log"
)

func CreateRoom(particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int, isCorrectWolf bool) string {

	room := structManege.CreateRoom{
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
