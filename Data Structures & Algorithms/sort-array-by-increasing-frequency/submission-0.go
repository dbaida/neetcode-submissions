import (
	"sort"
)

func frequencySort(nums []int) []int {
	numToFreqMap := make(map[int]int, 0)
	for _, num := range nums {
		numToFreqMap[num]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if numToFreqMap[nums[i]] != numToFreqMap[nums[j]] {
			return numToFreqMap[nums[i]] < numToFreqMap[nums[j]]
		}
		return nums[i] > nums[j]
	})

	return nums
}
