func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := abs(num)-1
		nums[idx] = -abs(nums[idx])
	}

	var result []int
	for i, appearance := range nums {
		if appearance > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}