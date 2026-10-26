func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	
	var candidate int
	
	HaystackLoop:
		for i := 0; i < len(haystack); i++ {
			if len(haystack)-i < len(needle) {
				return -1
			}
			if haystack[i] != needle[0] {
				continue
			}
			candidate = i
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					candidate = -1
					continue HaystackLoop
				}
			}
			return candidate
		}

	return -1
}
