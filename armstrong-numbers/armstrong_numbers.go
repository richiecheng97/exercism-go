package armstrong

import (
	"fmt"
	"math"
)

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}
	if n < 10 {
		return true
	}

	ll := len(fmt.Sprintf("%d", n))
	var num int
	tmp := n
	for {
		if tmp == 0 {
			break
		}
		num += int(math.Pow(float64(tmp%10), float64(ll)))
		tmp /= 10
	}
	return num == n
}
