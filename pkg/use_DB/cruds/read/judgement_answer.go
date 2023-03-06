package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func JudgementAnswer(roomId string) bool {
	ctx, client, _err := connectDB.ConnectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}
	answer := iter.Data()["IsCorrect"].(bool)
	defer client.Close()
	return answer

}
