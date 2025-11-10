package perfect

import "errors"

const (
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
	ClassificationDeficient Classification = "deficient"
)

var ErrOnlyPositive = errors.New("only positive integers are allowed")

// Define the Classification type here.
type Classification string

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}

	var res int64
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			res += i
			if i != n/i {
				res += n / i
			}
		}
	}

	res -= n
	switch {
	case res == n:
		return ClassificationPerfect, nil
	case res > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
