package encode

import "fmt"

func RunLengthEncode(input string) string {
	runes := []rune(input)
	var result string
	var count int

	for i, r := range runes {
		if i == 0 || r == runes[i-1] {
			count++
		} else {
			result += buildRuneEncoding(runes[i-1], count)
			count = 1
		}
	}
	result += buildRuneEncoding(runes[len(runes)-1], count)

	return result
}

func buildRuneEncoding(r rune, count int) string {
	if count > 1 {
		return fmt.Sprintf("%v%v", count, string(r))
	} else {
		return string(r)
	}
}

func RunLengthDecode(input string) string {
	runes := []rune(input)
	var result string
	var count string

	for _, r := range runes {
		if r >= '0' && r <= '9' {
			count += string(r)
		} else {
			if count == "" {
				result += string(r)
			} else {
				var n int
				fmt.Sscanf(count, "%d", &n)
				for i := 0; i < n; i++ {
					result += string(r)
				}
				count = ""
			}
		}
	}

	return result
}
