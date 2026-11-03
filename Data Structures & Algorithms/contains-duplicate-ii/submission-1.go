func containsNearbyDuplicate(nums []int, k int) bool {
	valueToRecentIdx := make(map[int]int)
	for idx, num := range nums {
		if jdx, ok := valueToRecentIdx[num]; ok {
			if abs(idx - jdx) <= k {
				return true
			}
		}
		valueToRecentIdx[num] = idx
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
