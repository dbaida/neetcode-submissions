func uncommonFromSentences(s1 string, s2 string) []string {
	wordsAppearance := make(map[string]int8)
	trackAppearance(wordsAppearance, s1)
	trackAppearance(wordsAppearance, s2)

	var res []string
	for word, count := range wordsAppearance {
		if count == 1 {
			res = append(res, word)
		}
	}
	
	return res
}

func trackAppearance(wordsAppearance map[string]int8, s string) {
	words := strings.Split(s, " ")
	for _, word := range words {
		wordsAppearance[word]++
	}
}
