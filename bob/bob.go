// Package bob contains function which defiend range of Bob responses.
package bob

import "strings"

// Hey returns bob responses based on specific remarks.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isEmpty := remark == ""
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := isShout(remark)

	var response string
	switch {
	case isEmpty:
		response = "Fine. Be that way!"
	case isQuestion && isYell:
		response = "Calm down, I know what I'm doing!"
	case isQuestion:
		response = "Sure."
	case isYell:
		response = "Whoa, chill out!"
	default:
		response = "Whatever."
	}
	return response
}

func isShout(str string) bool {
	return strings.ToUpper(str) == str && strings.ToLower(str) != str
}
