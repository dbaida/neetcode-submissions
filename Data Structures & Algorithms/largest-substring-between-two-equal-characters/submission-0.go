func maxLengthBetweenEqualCharacters(s string) int {
    var positions = [26]int{}
	for i := 0; i < 26; i++ {
		positions[i] = -1
	}

	maxLength := -1
	for i := 0; i < len(s); i++ {
		char := s[i] - 'a'
		
		if positions[char] == -1 {
			positions[char] = i
			continue
		}
		distance := i - positions[char] -1
		if distance > maxLength {
			maxLength = distance
		}
	}

	return maxLength
}