package darts

import "math"

func Score(x, y float64) int {
	var count int
	score := math.Abs(x)*math.Abs(x) + math.Abs(y)*math.Abs(y)
	if score >= 0 && score <= 1 {
		count += 10
	} else if score > 1 && score <= 25 {
		count += 5
	} else if score > 25 && score <= 100 {
		count += 1
	}
	return count
}
