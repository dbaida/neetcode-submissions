func search(nums []int, target int) int {
	for idx, num := range nums {
		if num == target {
			return idx
		}
	}
	return -1
}
