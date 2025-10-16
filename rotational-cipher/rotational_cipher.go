package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey %= 26
	if shiftKey < 0 {
		shiftKey += 26
	}

	var res strings.Builder
	res.Grow(len(plain))
	for _, cc := range plain {
		if cc >= 'a' && cc <= 'z' {
			res.WriteRune('a' + (cc-'a'+rune(shiftKey))%26)
		} else if cc >= 'A' && cc <= 'Z' {
			res.WriteRune('A' + (cc-'A'+rune(shiftKey))%26)
		} else {
			res.WriteRune(cc)
		}
	}
	return res.String()
}
