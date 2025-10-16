package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i
	}
	return res * res
}

func SumOfSquares(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i * i
	}
	return res
}

func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n) - SumOfSquares(n))))
}
