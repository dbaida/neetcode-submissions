func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	patternToWordMap := make(map[byte]string)
	wordToPatternMap := make(map[string]byte)

	for i := 0; i < len(words); i++ {
		if i == 0 {
			patternToWordMap[pattern[i]] = words[i]
			wordToPatternMap[words[i]] = pattern[i]
			continue
		}
		wordUsed, isPatternUsed := patternToWordMap[pattern[i]]
		if isPatternUsed && wordUsed != words[i] {
			return false
		}
		patternUsed, isWordUsed := wordToPatternMap[words[i]]
		if isWordUsed && patternUsed != pattern[i] {
			return false
		}
		patternToWordMap[pattern[i]] = words[i]
		wordToPatternMap[words[i]] = pattern[i]
	}
	return true
}