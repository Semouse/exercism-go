package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.Trim(id, " ")
	if len(id) <= 1 {
		return false
	}

	var digits []int
	for _, r := range id {
		if !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return false
		} else if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}

	sum := 0
	factor := 2
	for i := len(digits) - 1; i >= 0; i-- {
		factor = 3 - factor
		value := factor * digits[i]
		if value > 9 {
			value -= 9
		}

		sum += value
	}

	return sum%10 == 0
}
