package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Number cannot be less or equal 0 and be more than 64")
	}

	return uint64(math.Pow(2.0, float64(number-1))), nil
}

func Total() uint64 {
	return math.MaxUint64
}
