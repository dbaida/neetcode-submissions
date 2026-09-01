func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	first, last := nums[0], nums[len(nums)-1]

	if first > target {
		return 0
	}
	if last == target {
		return len(nums)-1
	}
	if last < target {
		return len(nums)
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] < target && nums[i+1] > target {
			return i+1
		}
	}
	return 0
}
