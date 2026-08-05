func majorityElement(nums []int) int {
	counters := make(map[int]int, 0)

    for _, num := range nums {
		counters[num]++
		count := counters[num]

		if count > len(nums) / 2 {
			return num
		}
	}

	return 0
}
