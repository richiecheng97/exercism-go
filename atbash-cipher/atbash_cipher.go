package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	var res strings.Builder
	var i = 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			res.WriteRune('a' + 'z' - r)
			i++
		} else if unicode.IsDigit(r) {
			res.WriteRune(r)
			i++
		} else {
			continue
		}
		if i%5 == 0 {
			res.WriteRune(' ')
		}
	}

	return strings.TrimSuffix(res.String(), " ")
}
