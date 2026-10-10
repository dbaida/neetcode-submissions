func intersection(nums1 []int, nums2 []int) []int {
	appeared := make(map[int]int8)
    for _, num1 := range nums1 {
		if _, exists := appeared[num1]; exists {
			continue
		}
		appeared[num1] = 1
	}

	var unique []int
	for _, num2 := range nums2 {
		if count, exists := appeared[num2]; !exists || count > 1 {
			continue
		}
		appeared[num2]++
		unique = append(unique, num2)
	}

	return unique
}
