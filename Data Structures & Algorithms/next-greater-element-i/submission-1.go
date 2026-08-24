func nextGreaterElement(nums1 []int, nums2 []int) []int {
	valToNextGreaterMap := make(map[int]int, len(nums2))

	for i := 0; i < len(nums2); i++ {
		nextGreater := -1

		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				nextGreater = nums2[j]
				break
			}
		}
		valToNextGreaterMap[nums2[i]] = nextGreater
	}

	res := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res = append(res, valToNextGreaterMap[nums1[i]])
	}
	return res
}
