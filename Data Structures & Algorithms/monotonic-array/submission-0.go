func isMonotonic(nums []int) bool {
    var isIncreasing, isDecreasing bool

	for i := 0; i < len(nums); i++ {
		if i != 0 && !isIncreasing && !isDecreasing {
			if nums[i] > nums[i-1] {
				isIncreasing = true
				continue
			}
			if nums[i] < nums[i-1] {
				isDecreasing = true
				continue
			}
		}

		if isIncreasing && nums[i] < nums[i-1] {
			return false
		}
		if isDecreasing && nums[i] > nums[i-1] {
			return false
		}
	}

	return true
}