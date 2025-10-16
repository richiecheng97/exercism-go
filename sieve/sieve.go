package sieve

func Sieve(limit int) []int {
	var res []int
	for i := 1; i <= limit; i++ {
		if IsPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// 检查从5开始的6k±1形式的数
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
