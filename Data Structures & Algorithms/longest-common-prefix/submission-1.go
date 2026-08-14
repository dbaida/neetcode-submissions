func longestCommonPrefix(strs []string) string {
	var prefix []byte

	firstWord := strs[0]

	for i := 0; i < len(firstWord); i++ {
		curr := firstWord[i]

		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != curr {
				return string(prefix)
			}
		}
		prefix = append(prefix, curr)
	}

	return string(prefix)
}
