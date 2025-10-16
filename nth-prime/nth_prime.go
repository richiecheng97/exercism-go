package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	primes := []int{}
	num := 2

	for len(primes) < n {
		if IsPrime(num) {
			primes = append(primes, num)
		}
		num++
	}

	return primes[n-1], nil
}

// IsPrime 判断一个数是否为素数
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
