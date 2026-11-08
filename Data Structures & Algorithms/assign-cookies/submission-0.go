func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var result int
	for i := 0; i < len(g) && len(s) > 0; i++ {
		for j := 0; j < len(s); j++ {
			if g[i] <= s[j] {
				s = append(s[:j], s[j+1:]...)
				result++
				break
			}
		}
	}

	return result
}
