package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	rr := []rune(isbn)
	if len(rr) != 10 {
		return false
	}
	var res int
	for i, v := range rr {
		if v == 'X' && i == len(rr)-1 {
			res += (10 - i) * 10
			continue
		}
		if v < '0' || v > '9' {
			return false
		}

		a, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		res += (10 - i) * a
	}

	return res%11 == 0
}
