func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var str1, str2 string
	for i := 0; i < len(word1); i++ {
		str1 += word1[i]
	}
	for j := 0; j < len(word2); j++ {
		str2 += word2[j]
	}

	return str1 == str2
}
