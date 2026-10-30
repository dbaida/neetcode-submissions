func finalPrices(prices []int) []int {
	l := len(prices)
	for i := 0; i < l-1; i++ {
		for j := i+1; j < l; j++ {
			if prices[j] > prices[i] {
				continue
			}
			prices[i] -= prices[j]
			break
		}
	}
	return prices
}
