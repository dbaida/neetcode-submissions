func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if (prices[j] - prices[i]) > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}
