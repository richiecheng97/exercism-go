package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subjectMap := genMap(subject)
	var result []string
	for _, str := range candidates {
		tmp := genMap(str)
		if reflect.DeepEqual(subjectMap, tmp) && subject != str && len(subject) == len(str) && !strings.EqualFold(subject, str) {
			result = append(result, str)
		}
	}
	return result
}

func genMap(str string) map[rune]int {
	str = strings.ToLower(str)
	mp := make(map[rune]int)
	for _, r := range str {
		// if !unicode.IsLetter(r) {
		// 	continue
		// }
		mp[r]++
	}
	return mp
}
