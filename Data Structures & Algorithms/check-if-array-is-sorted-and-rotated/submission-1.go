func check(nums []int) bool {
    var descCount int
	var prev int

	for i := 0; i < len(nums); i++ {
		isAsc := nums[i] >= prev
		prev = nums[i]

		if !isAsc {
			descCount++
			if descCount > 1 {
				return false
			}
		}
	}
	return descCount == 0 || nums[0] >= nums[len(nums)-1]
}