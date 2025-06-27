package reverse

func Reverse(input string) string {
	letters := []rune(input)
	for i, j := 0, len(letters)-1; i < len(letters)/2; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	return string(letters)
}
