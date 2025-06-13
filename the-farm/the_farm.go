package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	message string
	count   int
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.count, e.message)
}

// DivideFood returns amount of fodder per cow.
func DivideFood(fodderCalculator FodderCalculator, count int) (float64, error) {
	amount, err := fodderCalculator.FodderAmount(count)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amount / float64(count) * factor, nil
}

// ValidateInputAndDivideFood checks cow count and return amount of fodder per cow.
// If count <= 0 error is returned.
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	amount, err := DivideFood(fodderCalculator, count)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

// ValidateNumberOfCows return nil if count > 0 otherwise custom error is returned.
func ValidateNumberOfCows(count int) error {
	if count < 0 {
		return &InvalidCowsError{
			message: "there are no negative cows",
			count:   count,
		}
	} else if count == 0 {
		return &InvalidCowsError{
			message: "no cows don't need food",
			count:   count,
		}
	}
	return nil
}
