func findErrorNums(nums []int) []int {
	n := len(nums)
    freq := make([]int, n+1)
	for _, num := range nums {
		freq[num]++
	}

	res := []int{0, 0}
	var foundDup, foundMis bool
	for i := 1; i <= n; i++ {
		if freq[i] == 0 {
			res[1] = i
			if foundDup {
				return res
			}
			foundMis = true
		}
		if freq[i] == 2 {
			res[0] = i
			if foundMis {
				return res
			}
			foundDup = true
		}
	}
	return res
}