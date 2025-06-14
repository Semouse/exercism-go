package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var validLogRegExp = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var separatorRegExp = regexp.MustCompile(`<[*~=-]*?>`)
var passwordRegExp = regexp.MustCompile(`(?i)"*.password*."`)
var endOfLineRegExp = regexp.MustCompile(`end-of-line\d*`)
var userNameRegExp = regexp.MustCompile(`User\s+(\w+)`)

// Returns true if line start with any of: [TRC], [DBG], [INF], [WRN], [ERR], [FTL].
func IsValidLine(text string) bool {
	return validLogRegExp.MatchString(text)
}

// Returns split of string separated by separatorRegExp.
func SplitLogLine(text string) []string {
	return separatorRegExp.Split(text, -1)
}

// Returns number of matches for word password case incentive and in quotes.
func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if passwordRegExp.MatchString(line) {
			count++
		}
	}

	return count
}

// Removes end of the line text.nd remove the end-of-line text and return a "clean" string.
func RemoveEndOfLineText(text string) string {
	return endOfLineRegExp.ReplaceAllString(text, "")
}

// Append [USR] username if username is found.
func TagWithUserName(lines []string) []string {
	var result []string

	for _, line := range lines {
		matches := userNameRegExp.FindStringSubmatch(line)
		if matches != nil {
			line = fmt.Sprintf("[USR] %s %s", matches[1], line)
		}
		result = append(result, line)
	}

	return result
}
