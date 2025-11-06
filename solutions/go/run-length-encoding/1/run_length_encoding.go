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
			if count > 1 {
				result += fmt.Sprintf("%v%v", count, string(runes[i-1]))
			} else {
				result += string(runes[i-1])
			}
			count = 1
		}
	}
	if count > 1 {
		result += fmt.Sprintf("%v%v", count, string(runes[len(runes)-1]))
	} else if len(runes) > 0 {
		result += string(runes[len(runes)-1])
	}

	return result
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
