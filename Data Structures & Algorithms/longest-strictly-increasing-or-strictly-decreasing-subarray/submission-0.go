func longestMonotonicSubarray(nums []int) int {
    inc, dec, res := 1, 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			inc++
			dec = 1
		} else if nums[i-1] > nums[i] {
			dec++
			inc = 1
		} else {
			inc, dec = 1, 1
		}
		if inc > res {
			res = inc
		} else if dec > res {
			res = dec
		}
	}

	return res
}