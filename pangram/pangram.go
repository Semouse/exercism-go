package pangram

import "strings"

func IsPangram(input string) bool {
	var letters [26]int
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			letters[letter-'a']++
		}
	}

	for _, value := range letters {
		if value == 0 {
			return false
		}

	}

	return true

}
