func validWordAbbreviation(word string, abbr string) bool {
	var charsIncluded int
	var wordIndex int

	for i := 0; i < len(abbr); i++ {
		if wordIndex >= len(word) {
			return false
		}
		char := abbr[i]
		if isLetter(char) {
			if char != word[wordIndex] {
				return false
			}
			charsIncluded++
			wordIndex++
			continue
		}
		if char == '0' {
			return false
		}
		numberChars := make([]byte, 0)
		numberChars = append(numberChars, char)
		for j := i + 1; j < len(abbr); j++ {
			if isLetter(abbr[j]) {
				break
			}
			numberChars = append(numberChars, abbr[j])
			i++
		}
		number, _ := strconv.Atoi(string(numberChars))
		charsIncluded += number
		wordIndex += number
	}

	return charsIncluded == len(word)
}

func isLetter(b byte) bool {
	return b - 'a' >= 0 && b - 'a' <= 26
}
