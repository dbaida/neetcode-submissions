func relativeSortArray(arr1 []int, arr2 []int) []int {
	existingItems := make(map[int]int)
	for _, v2 := range arr2 {
		existingItems[v2] = 0
	}

	missingItems := make([]int, 0)
	for _, v1 := range arr1 {
		if _, exists := existingItems[v1]; exists {
			existingItems[v1]++
		} else {
			missingItems = append(missingItems, v1)
		}
	}

	res := make([]int, 0, len(arr1) + len(arr2))
	for _, v2 := range arr2 {
		count, _ := existingItems[v2]
		for i := 0; i < count; i++ {
			res = append(res, v2)
		}
	}
	sort.Ints(missingItems)

	return append(res, missingItems...)
}