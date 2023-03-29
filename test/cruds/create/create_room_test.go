package createDBTest

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"gamename-back-end/test/db"

	createDB "gamename-back-end/pkg/cruds/create"
)

func Test_CreateRoom(t *testing.T) {
	ctx, client, err := db.ConnectDBForTest()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}

	// NOTE: コレクションがない場合には、その場で作成されるため、このテストの範囲内では InitializeDatabase は不要
	db.InitializeDatabase(ctx, client)
	defer client.Close()

	// NOTE: DB の接続周りのテストを一応しておく
	t.Run("ctxが存在する", func(t *testing.T) {
		if got := ctx; got == nil {
			t.Errorf("got = %v, want %v", got, nil)
		}
	})
	t.Run("clientが存在する", func(t *testing.T) {
		if got := client; got == nil {
			t.Errorf("client = %v, want %v", got, nil)
		}
	})

	// ランダムな ID の作成
	const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())

	roomIdByte1 := make([]byte, 20)
	for i := range roomIdByte1 {
		roomIdByte1[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	roomIdByte2 := make([]byte, 20)
	for i := range roomIdByte2 {
		roomIdByte2[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	roomId1 := string(roomIdByte1)
	roomId2 := string(roomIdByte2)

	tests := []struct {
		name       string
		isWolfMode bool
		roomId     string
	}{
		{name: "wolfModeがfalseの場合", isWolfMode: false, roomId: roomId1},
		{name: "wolfModeがtrueの場合", isWolfMode: true, roomId: roomId2},
	}
	for _, tt := range tests {
		t.Run(tt.name+"Roomが作成できる", func(t *testing.T) {
			var id = createDB.CreateRoom(ctx, client, 0, "theme", 0, 0, tt.isWolfMode, false, 0, true, tt.roomId)
			if got := id; got != tt.roomId {
				t.Errorf("get id is empty, want %v", got)
			}
		})
	}

	db.DeleteCollection(ctx, client, 100, "Room")
}
