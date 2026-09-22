func countBits(n int) []int {
	result := make([]int, 0, n)
	result = append(result, 0)

	for i := 1; i <= n; i++ {
		result = append(result, bits.OnesCount(uint(i)))
	}

	return result
}
