package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	// var total uint64
	// for i := 1; i < 65; i++ {
	// 	t, _ := Square(i)
	// 	total += t
	// }
	// return total
	return 18446744073709551615
}
