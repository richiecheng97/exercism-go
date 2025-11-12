package piglatin

import "strings"

func Sentence(sentence string) string {
	if len(sentence) == 0 {
		return sentence
	}

	if sentence[0] == 'a' || sentence[0] == 'e' || sentence[0] == 'i' || sentence[0] == 'o' || sentence[0] == 'u' {
		return sentence + "ay"
	} else if len(sentence) >= 2 && (sentence[:2] == "xr" || sentence[:2] == "yt") {
		return sentence + "ay"
	} else {
		indexQu := strings.Index(sentence, "qu")
		indexY := strings.Index(sentence, "y")

		var tmp string
		for _, v := range sentence {
			if sentence[0] != 'a' && sentence[0] != 'e' && sentence[0] != 'i' && sentence[0] != 'o' && sentence[0] != 'u' {
				tmp += string(v)
			} else {
				break
			}
		}
		if indexQu >= 0 && indexQu == len(tmp) {
			return sentence[len(tmp)+2:] + tmp + "qu" + "ay"
		} else if indexY >= 0 && indexY == len(tmp) {
			return sentence[len(tmp):] + tmp + "ay"
		}
		return sentence + "ay"
	}
}
