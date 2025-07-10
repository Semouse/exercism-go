package isbn

func getDigit(letter rune) (digit int, ok bool) {
	if letter < '0' || letter > '9' {
		return digit, ok
	}
	return int(letter - '0'), true
}

func IsValidISBN(isbn string) bool {
	count := 10
	sum := 0

	for _, letter := range isbn {
		if digit, ok := getDigit(letter); ok {
			sum += (digit * count)
			count--
		} else if count == 1 && letter == 'X' {
			sum += 10
			count--
		} else if letter != '-' {
			return false
		}
	}

	return count == 0 && sum%11 == 0
}
