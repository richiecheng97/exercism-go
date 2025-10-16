package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(ff FodderCalculator, cows int) (float64, error) {
	total, err := ff.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	res, err := ff.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return total * res / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(ff FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(ff, cows)
	}
	return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", cows)
	}
	if cows == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", cows)
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
