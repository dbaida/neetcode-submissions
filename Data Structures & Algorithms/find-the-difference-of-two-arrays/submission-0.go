func findDifference(nums1 []int, nums2 []int) [][]int {
	hashSet1 := make(map[int]bool)
	for _, num1 := range nums1 {
		hashSet1[num1] = true
	}
	answer2 := make([]int, 0)
	hashSet2 := make(map[int]bool)
	for _, num2 := range nums2 {
		if !hashSet2[num2] && !hashSet1[num2] {
			answer2 = append(answer2, num2)
		}
		hashSet2[num2] = true
	}
	answer1 := make([]int, 0)
	for num1, _ := range hashSet1 {
		if !hashSet2[num1] {
			answer1 = append(answer1, num1)
		}
	}
	return [][]int{answer1, answer2}
}
