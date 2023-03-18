// ---------------

// テストの実行方法

// ---------------
// go test -v ./...

package test_temp

import "testing"

// 関数名は Test_### にする必要がある
func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}

	// 1個テストする
	t.Run("正常を1個テスト", func(t *testing.T) {
		if got := add(1, 2); got != 3 {
			t.Errorf("add() = %v, want %v", got, 3)
		}
	})

	// 複数テストする
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "正常-1", args: args{a: 1, b: 2}, want: 3},
		{name: "正常-2", args: args{a: -1, b: -2}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
