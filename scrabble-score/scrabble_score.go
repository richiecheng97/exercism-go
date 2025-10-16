package scrabble

import "strings"

func Score(word string) int {
	var res int
	for _, v := range strings.ToUpper(word) {
		switch v {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res += 1
		case 'D', 'G':
			res += 2
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'K':
			res += 5
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		}
	}
	return res
}
