package raindrops

import "strconv"

var numbers = [...]int{3, 5, 7}
var drops = [...]string{"Pling", "Plang", "Plong"}

// Convert number to a string by following rules.
// If a given number:
// is divisible by 3, add "Pling" to the result.
// is divisible by 5, add "Plang" to the result.
// is divisible by 7, add "Plong" to the result.
// is not divisible by 3, 5, or 7, the result should be the number as a string.
func Convert(number int) string {
	var result string
	for i, n := range numbers {
		if number%n == 0 {
			result += drops[i]
		}
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
