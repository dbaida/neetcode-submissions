func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
	seen := make(map[int]bool)
	var repeated int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if seen[grid[i][j]] {
				repeated = grid[i][j]
			} else {
				seen[grid[i][j]] = true
			}
		}
	}

	var missing int
	for val := 1; val <= n*n; val++ {
		if !seen[val] {
			missing = val
			break
		}
	}

	return []int{repeated, missing}
}