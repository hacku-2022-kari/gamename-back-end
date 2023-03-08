package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"

	"cloud.google.com/go/firestore"
)

// 0(平和村,true),1(平和村,false),2(人狼村,true),3(人狼村,false)
func JudgementWolf(roomId string, playerId string) int {
	ctx, client, err := connectDB.ConnectDB()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	var branch = []bool{true, true}

	roomRef := client.Collection("Room").Doc(roomId)

	roomIter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}
	if roomIter.Data()["IsExitWolf"].(bool) == true {
		branch[0] = false
	}

	_, err = client.Collection("Room").Doc(roomId).Update(ctx, []firestore.Update{
		{Path: "Step", Value: 10},
	})
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}

	playerIter, err := client.Collection("Player").Doc(playerId).Get(ctx)

	defer client.Close()

	if err != nil {
		if branch[0] == false {
			_, err = roomRef.Update(ctx, []firestore.Update{
				{Path: "IsCorrectWolf", Value: false},
			})
			if err != nil {
				log.Printf("error getting Room documents: %v\n", err)
			}
			_, err = roomRef.Update(ctx, []firestore.Update{
				{Path: "Result", Value: 4},
			})
			return 4
		} else {
			_, err = roomRef.Update(ctx, []firestore.Update{
				{Path: "Result", Value: 1},
			})
			return 1
		}
	}
	if playerIter.Data()["Wolf"].(bool) == true {
		if branch[0] == true {
			branch[1] = false
		}
	} else {
		if branch[0] == true {
			branch[1] = false
		}
		if branch[0] == false {
			branch[1] = false
		}
	}

	if branch[0] == true && branch[1] == true {
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "Result", Value: 1},
		})
		return 1
	} else if branch[0] == true && branch[1] == false {
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "IsCorrectWolf", Value: false},
		})
		if err != nil {
			log.Printf("error getting Room documents: %v\n", err)
		}
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "Result", Value: 2},
		})
		return 2
	} else if branch[0] == false && branch[1] == true {
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "Result", Value: 3},
		})
		return 3
	} else {
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "IsCorrectWolf", Value: false},
		})
		if err != nil {
			log.Printf("error getting Room documents: %v\n", err)
		}
		_, err = roomRef.Update(ctx, []firestore.Update{
			{Path: "Result", Value: 4},
		})
		return 4
	}

}
