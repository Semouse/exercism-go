package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if word is isogram.
// Isograms a word or phrase without a repeating letter.
func IsIsogram(word string) bool {
	var letters []rune
	for _, letter := range strings.ToLower(word) {
		for _, value := range letters {
			if value == letter && unicode.IsLetter(value) {
				return false
			}
		}

		letters = append(letters, letter)
	}

	return true
}
