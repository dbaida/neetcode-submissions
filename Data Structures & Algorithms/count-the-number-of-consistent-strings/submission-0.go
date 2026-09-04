func countConsistentStrings(allowed string, words []string) int {
    chars := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		chars[allowed[i]] = true
	}

	var result int
	for _, word := range words {
		consistent := true
		for j := 0; j < len(word); j++ {
			if _, ok := chars[word[j]]; !ok {
				consistent = false
				break
			}
		}
		if consistent {
			result++
		}
	}
	return result
}