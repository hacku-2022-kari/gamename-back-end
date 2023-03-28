package createDB

import (
	"context"
	"fmt"
	types "gamename-back-end/pkg/types"

	"cloud.google.com/go/firestore"
)

func CreateRoom(ctx context.Context, client *firestore.Client, particNum int, theme string, phase int, step int, wolfMode bool, isExitWolf bool, peaceVote int, isCorrectWolf bool, roomId string) string {

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

	_, err := client.Collection("Room").Doc(roomId).Set(ctx, room)
	if err != nil {
		fmt.Println("An error has occurred", err)
		// NOTE: ID の作成に失敗した場合には空文字を返す
		return ""
	}

	return roomId
}
