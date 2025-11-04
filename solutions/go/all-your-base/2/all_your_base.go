package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	firstNonZero := 0
	for firstNonZero < len(inputDigits) && inputDigits[firstNonZero] == 0 {
		firstNonZero++
	}

	if firstNonZero == len(inputDigits) {
		return []int{0}, nil
	}

	trimmed := inputDigits[firstNonZero:]
	value := 0
	for _, d := range trimmed {
		value = value*inputBase + d
	}

	var out []int
	for value > 0 {
		out = append(out, value%outputBase)
		value /= outputBase
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out, nil
}
