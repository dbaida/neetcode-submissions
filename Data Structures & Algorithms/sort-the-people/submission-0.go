func sortPeople(names []string, heights []int) []string {
	heightToNameMap := make(map[int]string, len(names))
	for i := 0; i < len(names); i++ {
		heightToNameMap[heights[i]] = names[i]
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[j] > heights[i]
	})

	sortedNames := make([]string, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		sortedNames = append(sortedNames, heightToNameMap[heights[i]])
	}

	return sortedNames
}