func isArraySpecial(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	isPrevEven := nums[0] % 2 == 0
	for i := 1; i < n; i++ {
		if isPrevEven {
			if nums[i] % 2 == 0 {
				return false
			}
			isPrevEven = false
		} else {
			if nums[i] % 2 != 0 {
				return false
			}
			isPrevEven = true
		}
	}
	return true
}