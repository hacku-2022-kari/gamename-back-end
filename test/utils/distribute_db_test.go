package utilsTest

import (
	utils "gamename-back-end/pkg/utils"
	"testing"
)

func Test_DistributeDB(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "頭文字が数字の場合", input: "123abc", expected: "0"},
		{name: "頭文字がa-jの場合", input: "abc123", expected: "1"},
		{name: "頭文字がk-tの場合", input: "klm123", expected: "2"},
		{name: "頭文字がu-z,A-Dの場合", input: "uvw123", expected: "3"},
		{name: "頭文字がE-Nの場合", input: "EPQ123", expected: "4"},
		{name: "頭文字がO-Zの場合", input: "XYZ123", expected: "5"},
		{name: "入力が空文字列の場合", input: "", expected: ""},
		{name: "入力が有効な値ではない場合", input: "*", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.DistributeDB(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, result)
			}
		})
	}
}
