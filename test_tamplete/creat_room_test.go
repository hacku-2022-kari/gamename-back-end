package test_temp

import (
	useDB "gamename-back-end/src/use_DB"
	"testing"
)

type room struct {
	Password string
	PaticNum int
	Theme    string
	Phase    int
	Step     int
}

func TestCreateRoom(t *testing.T) {
	t.Run("正常を1個テスト", func(t *testing.T) {
		if got := useDB.CreateRoom("pass", 0, "no-theme", 1, 0); got != "" {
			t.Errorf("CreatRoom = %v, want NotNUll", got)
		}
	})

	tests := []struct {
		name string
		args room
	}{
		{name: "正常-1", args: room{Password: "test1", PaticNum: 0, Theme: "test", Phase: 0, Step: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := useDB.CreateRoom(tt.args.Password, tt.args.PaticNum, tt.args.Theme, tt.args.Phase, tt.args.Step); got != "" {
				t.Errorf("CreateRoom = %v, want NotNUll", got)
			}
		})
	}
}
