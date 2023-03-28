package readDB

import (
	"encoding/json"
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

// TODO: 構造体をまとめたところに移す
type RoleWolf struct {
	Role int  `json:"role"`
	Wolf bool `json:"wolf"`
}

// NOTE: roomId は ConnectDB の引数を増やすタイミングで使用するため一旦未使用で良い
func GetRoleWolf(playerId string, roomId string) RoleWolf {
	ctx, client, err := connectDB.ConnectDB(roomId)
	if err != nil {
		log.Println("error getting Player document: \n", err)

	}
	iter, err := client.Collection("Player").Doc(playerId).Get(ctx)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}

	bytes, _ := json.Marshal(iter.Data()["Role"])
	var roleInt int64
	err = json.Unmarshal(bytes, &roleInt)
	if err != nil {
		log.Printf("error getting Room documents: %v\n", err)
	}

	var roleWolf RoleWolf
	roleWolf.Role = int(roleInt)
	roleWolf.Wolf = iter.Data()["Wolf"].(bool)
	defer client.Close()

	return roleWolf
}
