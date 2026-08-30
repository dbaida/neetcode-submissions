func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)
	for _, char := range magazine {
		letters[char]++
	}
	for _, char := range ransomNote {
		if count, ok := letters[char]; !ok || count == 0 {
			return false
		}
		letters[char]--
	}
	return true
}
