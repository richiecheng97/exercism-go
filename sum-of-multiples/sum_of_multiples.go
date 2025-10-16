package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var mp = make(map[int]bool, 0)
	for _, v := range divisors {
		var i = 1
		for {
			if i*v >= limit || i*v <= 0 {
				break
			}
			mp[i*v] = true
			i++
		}
	}

	var total int
	for k := range mp {
		total += k
	}
	return total
}
