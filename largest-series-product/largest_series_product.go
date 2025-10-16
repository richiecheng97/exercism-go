package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}

	rr := []rune(digits)
	var res int
	for i := 0; i <= len(digits)-span; i++ {
		sum := 1
		for j := 0; j < span; j++ {
			if !unicode.IsDigit(rr[i+j]) {
				return 0, errors.New("digits input must only contain digits")
			}
			a, _ := strconv.Atoi(string(rr[i+j]))
			sum *= a
		}
		if sum > res {
			res = sum
		}
	}
	return int64(res), nil
}
