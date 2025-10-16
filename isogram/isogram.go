package isogram

import "strings"

func IsIsogram(word string) bool {
	mp := make(map[rune]bool, len(word))
	for _, v := range strings.ToUpper(word) {
		if v >= 'A' && v <= 'Z' {
			if _, ok := mp[v]; ok {
				return false
			}
			mp[v] = true
		}
	}
	return true
}
