func countCharacters(words []string, chars string) int {
	canFormWordFromChars := func(word string, counts [26]int) bool {
		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'
			if counts[idx] == 0 {
				return false
			}
			counts[idx]--
		}
		return true
	}

    var availableChars [26]int
	for i := 0; i < len(chars); i++ {
		availableChars[chars[i]-'a']++
	}
	
	var result int
	for _, word := range words {
		// 3rd case contains hidden char which brakes implementation
		// that's why we need to sanitize input
		if canFormWordFromChars(strings.TrimSpace(word), availableChars) {
			result += len(word)
		}
	}

	return result
}