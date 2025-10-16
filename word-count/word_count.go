package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	frequency := make(Frequency)
	var word strings.Builder

	for i, r := range phrase {
		isAlp := unicode.IsLetter(r) || unicode.IsDigit(r)
		isApo := r == '\''

		if isAlp {
			word.WriteRune(r)
		} else if isApo {
			prevIsAlp := i > 0 && (unicode.IsLetter(rune(phrase[i-1])) || unicode.IsDigit(rune(phrase[i-1])))
			nextIsAlp := i+1 < len(phrase) && (unicode.IsLetter(rune(phrase[i+1])) || unicode.IsDigit(rune(phrase[i+1])))

			if prevIsAlp && nextIsAlp {
				word.WriteRune(r)
			} else {
				toFrequency(frequency, &word)
			}
		} else {
			toFrequency(frequency, &word)
		}
	}
	toFrequency(frequency, &word)

	return frequency
}

func toFrequency(frequency Frequency, word *strings.Builder) {
	res := word.String()
	if res != "" {
		frequency[res]++
	}
	word.Reset()
}
