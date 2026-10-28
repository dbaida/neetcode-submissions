func arrangeCoins(n int) int {
	var requiredCoins, completeRows int

	for {
		requiredCoins++
		if n < requiredCoins {
			break
		}
		n -= requiredCoins
		completeRows++
	}

	return completeRows
}
