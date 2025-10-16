package prime

func Factors(n int64) []int64 {
	if n <= 1 {
		return []int64{}
	}

	factors := []int64{}

	// 检查因子2
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// 检查奇数因子，从3开始，步长为2
	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// 如果temp是大于2的质数
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}
