package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
	"math/rand"
	"time"
)

func GetRandomTheme() string {
	ctx, client, err := connectDB.ConnectDB()

	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	docs, err := client.Collection("Random_theme").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(docs))

	docRef := docs[idx].Ref

	iter, err := client.Collection("Random_theme").Doc(docRef.ID).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	theme := iter.Data()["theme"].(string)
	defer client.Close()
	return theme
}
