package useDB

import (
	"log"
)

// 0(平和村,true),1(平和村,false),2(人狼村,true),3(人狼村,false)
func JudgementWolf(roomId string, playerId string) int {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	var branch = []bool{true, true}
	roomIter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}
	if roomIter.Data()["IsExitWolf"].(bool) == true {
		branch[0] = false
	}
	playerIter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	defer client.Close()
	if err != nil {
		if branch[0] == false {
			return 4
		} else {
			return 1
		}
	}
	if playerIter.Data()["IsWolf"].(bool) == true {
		if branch[0] == true {
			branch[1] = false
		}
	} else {
		if branch[0] == false {
			branch[1] = false
		}
	}

	if branch[0] == true && branch[1] == true {
		return 1
	} else if branch[0] == true && branch[1] == false {
		return 2
	} else if branch[0] == false && branch[1] == true {
		return 3
	} else {
		return 4
	}

}

// $body = @{
// 	roomId = "WgBySaSIBvs92OsDdd4i"
//     playerId = "W8fAxy12FB8fGF9vysxy"
// } | ConvertTo-Json
// Invoke-RestMethod -Method POST -Uri http://localhost:1323/judgement-wolf -Body $body -ContentType "application/json"
