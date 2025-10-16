package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden struct {
	mp map[string][]string
}

var flag = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	children = RemoveEmptyAndDuplicatesFast(children)
	sort.Strings(children)

	if len(strings.Split(diagram, "\n")) != 3 {
		return nil, errors.New("")
	}

	dd := []rune(strings.ReplaceAll(diagram, "\n", ""))
	ld := len(dd)
	lc := len(children)

	// fmt.Printf("%v\n", ld == 0)
	// fmt.Printf("%v\n", lc == 0)
	// fmt.Printf("%v\n", ld/4 != lc)
	// fmt.Printf("%v\n", ld%4 != 0)
	if ld == 0 || lc == 0 || ld/4 != lc || ld%4 != 0 {
		return nil, errors.New("")
	}

	var mp = make(map[string][]string, lc)
	var k = 0
	for j := 0; j < 2; j++ {
		for _, c := range children {
			rr1 := dd[k]
			k++
			rr2 := dd[k]
			k++
			if v, ok := flag[rr1]; ok {
				mp[c] = append(mp[c], v)
			} else {
				return nil, errors.New("")
			}
			if v, ok := flag[rr2]; ok {
				mp[c] = append(mp[c], v)
			} else {
				return nil, errors.New("")
			}
		}
	}

	return &Garden{
		mp: mp,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if v, ok := g.mp[child]; ok {
		return v, true
	}
	return nil, false
}

func RemoveEmptyAndDuplicatesFast(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[string]struct{}, len(slice))
	result := make([]string, 0, len(slice))

	for _, str := range slice {
		// 快速检查空字符串
		if str == "" {
			continue
		}

		// 去除首尾空白字符
		trimmed := strings.TrimLeft(strings.TrimRight(str, " \t\n\r"), " \t\n\r")
		if trimmed == "" {
			continue
		}

		// 去重检查
		if _, exists := seen[trimmed]; !exists {
			seen[trimmed] = struct{}{}
			result = append(result, trimmed)
		}
	}

	return result
}
