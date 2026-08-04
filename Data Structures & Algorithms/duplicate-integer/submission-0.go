func hasDuplicate(nums []int) bool {
	contains := make(map[int]bool, len(nums))

    for _, val := range nums {
		if ok := contains[val]; ok {
			return true
		}

		contains[val] = true
	}

	return false
}

