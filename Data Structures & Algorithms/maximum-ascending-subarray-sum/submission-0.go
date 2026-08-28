func maxAscendingSum(nums []int) int {
    var max, acc, prev int

	for _, num := range nums {
		if num <= prev {
			acc = 0
		}
		acc += num

		if acc > max {
			max = acc
		}
		prev = num
	}

	return max
}