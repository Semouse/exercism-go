package logs

// Stores specific log character bound to application.
var emmiters = map[rune]string{
	'\U0001F50D': "search",
	'\u2757':     "recommendation",
	'\u2600':     "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, ch := range log {
		application, present := emmiters[ch]
		if present {
			return application
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	letters := []rune(log)
	for i, letter := range letters {
		if oldRune == letter {
			letters[i] = newRune
		}
	}

	return string(letters)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
	for range log {
		count++
	}

	return count <= limit
}
