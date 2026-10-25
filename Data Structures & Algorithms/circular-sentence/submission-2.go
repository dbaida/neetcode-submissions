func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	var lastChar byte
	for _, word := range words {
		if lastChar != 0 && word[0] != lastChar {
			return false
		}
		lastChar = word[len(word)-1]
	}
	return true
}
