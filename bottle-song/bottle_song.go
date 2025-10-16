package bottlesong

import (
	"fmt"
	"strings"
)

var strs = []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Recite(startBottles, takeDown int) []string {
	var res []string
	for i := 0; i < takeDown; i++ {
		if startBottles == 1 {
			res = append(res, fmt.Sprintf("%s green bottle hanging on the wall,", strs[startBottles]))
			res = append(res, fmt.Sprintf("%s green bottle hanging on the wall,", strs[startBottles]))
		} else {
			res = append(res, fmt.Sprintf("%s green bottles hanging on the wall,", strs[startBottles]))
			res = append(res, fmt.Sprintf("%s green bottles hanging on the wall,", strs[startBottles]))
		}
		res = append(res, "And if one green bottle should accidentally fall,")
		if startBottles == 2 {
			res = append(res, fmt.Sprintf("There'll be %s green bottle hanging on the wall.", strings.ToLower(strs[startBottles-1])))
		} else {
			res = append(res, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(strs[startBottles-1])))
		}
		if i != takeDown-1 {
			res = append(res, "")
		}
		startBottles--
	}
	return res
}

/**
["One green bottles hanging on the wall," "One green bottles hanging on the wall," "And if one green bottle should accidentally fall," "There'll be no green bottles hanging on the wall."], want:
["One green bottle hanging on the wall," "One green bottle hanging on the wall," "And if one green bottle should accidentally fall," "There'll be no green bottles hanging on the wall."]

*/
