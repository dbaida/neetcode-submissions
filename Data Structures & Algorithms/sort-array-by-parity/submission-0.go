func sortArrayByParity(nums []int) []int {
	result := make([]int, 0, len(nums))
	odd := make([]int, 0)

	for _, num := range nums {
		if num % 2 == 0 {
			result = append(result, num)
		} else {
			odd = append(odd, num)
		}
	}

	return append(result, odd...)
}
