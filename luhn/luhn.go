package luhn

import "strings"

func Valid(id string) bool {
	id = strings.TrimSpace(id)
	if len(id) <= 1 {
		return false
	}

	var j = 1
	var total int
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue
		}
		if id[i] < '0' || id[i] > '9' {
			return false
		}

		num := int(id[i] - '0')
		if j%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		total += num
		j++
	}

	return total%10 == 0
}
