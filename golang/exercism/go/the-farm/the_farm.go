package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

// TODO: define the 'DivideFood' function

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fod, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	fat, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fod / float64(cows) * fat, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fc, cows)
	}
	return 0, errors.New("invalid number of cows")
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.cows, err.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
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
