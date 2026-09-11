import "slices"

func heightChecker(heights []int) int {
    expected := slices.Clone(heights)
	slices.Sort(expected)
	
	var result int
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			result++
		}
	}
	return result
}