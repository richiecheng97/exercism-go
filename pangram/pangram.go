package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToUpper(input)
	var mp = make(map[rune]bool, 0)
	for _, v := range input {
		if v >= 'A' && v <= 'Z' {
			mp[v] = true
		}
	}

	return len(mp) == 26
}
