func divideArray(nums []int) bool {
    pairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pairs[nums[i]]++
	}
	for _, count := range pairs {
		if count % 2 != 0 {
			return false
		}
	}
	return true
}