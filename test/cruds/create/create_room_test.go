package createDBTest

import (
	"log"
	"testing"

	"gamename-back-end/test/db"

	createDB "gamename-back-end/pkg/cruds/create"
)

func Test_CreateRoom(t *testing.T) {
	ctx, client, err := db.ConnectDBForTest()
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
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

	tests := []struct {
		name       string
		isWolfMode bool
	}{
		{name: "wolfModeがfalseの場合", isWolfMode: false},
		{name: "wolfModeがtrueの場合", isWolfMode: true},
	}
	for _, tt := range tests {
		t.Run(tt.name+"Roomが作成できる", func(t *testing.T) {
			var id = createDB.CreateRoom(ctx, client, 0, "theme", 0, 0, tt.isWolfMode, false, 0, true, "roomId")
			if got := id; got == "" {
				t.Errorf("get id is %v", id)
			}
		})
	}

	db.DeleteCollection(ctx, client, 100, "Room")
}
