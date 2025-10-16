package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var mp = make(map[string]int)
	for k, v := range in {
		for _, s := range v {
			s = strings.ToLower(s)
			mp[s] = k
		}
	}
	return mp
}
