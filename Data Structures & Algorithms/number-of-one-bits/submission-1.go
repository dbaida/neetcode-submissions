func hammingWeight(n int) int {
	if n == 0 {
		return 0
	}
	var result int
	var hasRemainder bool
	quotient := n

	for {
		hasRemainder = quotient % 2 != 0
		quotient = quotient / 2

		if quotient == 0 || hasRemainder {
			result++
			if quotient == 0 {
				break
			}
		}
	}

	return result
}
