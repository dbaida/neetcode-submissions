func firstUniqChar(s string) int {
	var chars [26]int
	for i := 0; i < len(chars); i++ {
		chars[i] = math.MaxInt
	}
	for k, v := range s {
		// ignore non-unique chars
		if chars[v - 'a'] == -1 {
			continue
		}
		// not touched previously
		if chars[v - 'a'] == math.MaxInt {
			chars[v - 'a'] = k
			continue
		}
		// mark as non-unique
		chars[v - 'a'] = -1
	}
	smallest := math.MaxInt
	for i := 0; i < len(chars); i++ {
		if chars[i] != -1 && smallest > chars[i] {
			smallest = chars[i]
		}
	}
	if smallest == math.MaxInt {
		return -1
	}
	return smallest
}
