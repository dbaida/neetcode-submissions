func removeDuplicates(nums []int) int {
	i, n := 1, len(nums)

	for i < n {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			n--
		} else {
			i++
		}
	}
	return len(nums)
}
