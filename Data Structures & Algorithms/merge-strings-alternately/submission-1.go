func mergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)
	
	maxLen := len2
	if len1 > len2 {
		maxLen = len1
	}

	merged := make([]byte, 0, len1+len2)
	for i := 0; i < maxLen; i++ {
		if i < len1 {
			merged = append(merged, word1[i])
		}
		if i < len2 {
			merged = append(merged, word2[i])
		}
	}
	return string(merged)
}
