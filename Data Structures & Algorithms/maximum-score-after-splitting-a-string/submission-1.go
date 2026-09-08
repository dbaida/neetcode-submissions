func maxScore(s string) int {
    var maxScore, leftScore, rightScore int
	for i := 0; i < len(s) - 1; i++ {
		leftScore, rightScore = 0, 0
		for l := i; l >= 0; l-- {
			if s[l] == '0' {
				leftScore++
			}
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == '1' {
				rightScore++
			}
		}
		if leftScore + rightScore > maxScore {
			maxScore = leftScore + rightScore
		}
	}
	return maxScore
}