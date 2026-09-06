func maxNumberOfBalloons(text string) int {
	baloonChars := map[rune]int{
		'b'-'a': 1,
		'a'-'a': 1,
		'l'-'a': 2,
		'o'-'a': 2,
		'n'-'a': 1,
	}

	var textChars [26]int
	for i := 0; i < len(text); i++ {
		textChars[text[i]-'a']++
	}

	maxNumber := math.MaxInt
	for char, requiredCount := range baloonChars {
		if requiredCount > textChars[char] {
			return 0
		}
		maxWords := textChars[char] / requiredCount
		if maxNumber > maxWords {
			maxNumber = maxWords
		}
	}

	return maxNumber
}