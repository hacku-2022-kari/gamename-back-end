package utils

func DistributeDB(str string) string {
	if len(str) == 0 {
		return ""
	}

	switch str[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return "0"
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j':
		return "1"
	case 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't':
		return "2"
	case 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D':
		return "3"
	case 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N':
		return "4"
	case 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return "5"
	default:
		return "" // ここに来ることはない
	}
}
