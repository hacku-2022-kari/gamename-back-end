package createDB

import (
	"context"
	types "gamename-back-end/pkg/types"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateRoom(ctx context.Context, client *firestore.Client, particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int, isCorrectWolf bool) string {

	room := types.CreateRoom{
		ParticNum:     particNum,
		Theme:         theme,
		Phase:         phase,
		Step:          step,
		IsModeWolf:    wolfMode,
		IsExitWolf:    isExitWolf,
		PeaceVote:     peaceVote,
		IsCorrectWolf: isCorrectWolf,
		Result:        1,
	}

	ref := client.Collection("Room").NewDoc()
	_, err := ref.Set(ctx, room)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		// NOTE: ID の作成に失敗した場合には空文字を返す
		return ""
	}

	return ref.ID
}
