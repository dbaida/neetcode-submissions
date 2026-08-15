func scoreOfString(s string) int {
    score := 0

    for i := 0; i < len(s) - 1; i++ {
        score += abs(int(s[i + 1]) - int(s[i]))
    }

    return score
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
