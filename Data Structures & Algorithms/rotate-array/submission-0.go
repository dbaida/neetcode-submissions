func rotate(nums []int, k int) {
	for ; k > 0; k-- {
		for i := 0; i < len(nums); i++ {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
		}
	}
}
