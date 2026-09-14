func findLucky(arr []int) int {
	counters := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		counters[arr[i]]++
	}

	largestNum := -1
	for num, freq := range counters {
		if num == freq && num > largestNum {
			largestNum = num
		}
	}
	return largestNum
}
