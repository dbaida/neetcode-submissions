func isPalindrome(s string) bool {
	transformed := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, strings.ToLower(s))

	reversed := make([]byte, 0, len(transformed))

	for i := len(transformed) - 1; i >= 0; i-- {
		reversed = append(reversed, transformed[i])
	}
	
	return transformed == string(reversed)
}
