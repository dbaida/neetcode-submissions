func commonChars(words []string) []string {
    commonChars := make(map[rune]int, 0)
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordChars := make(map[rune]int, 0)
		for _, char := range word {
			wordChars[char]++
		}
		if i == 0 {
			for k, v := range wordChars {
				commonChars[k] = v
			}
		} else {
			for commonChar, commonCharCount := range commonChars {
				if wordCharCount, exists := wordChars[commonChar]; exists {
					if wordCharCount < commonCharCount {
						commonChars[commonChar] = wordCharCount
					}
				} else {
					delete(commonChars, commonChar)
				}
			}
		}
	}
	result := make([]string, 0)
	for char, repeats := range commonChars {
		if repeats == 1 {
			result = append(result, string(char))
		} else {
			for i := 1; i <= repeats; i++ {
				result = append(result, string(char))
			}
		}
	}
	return result
}