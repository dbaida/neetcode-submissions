func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMapS := make(map[rune]int)
	hashMapT := make(map[rune]int)

	for i, char := range s {
		hashMapS[char]++
		hashMapT[rune(t[i])]++
	}

	for char, count := range hashMapS {
		if hashMapT[char] != count {
			return false
		}
	}

	return true
}
