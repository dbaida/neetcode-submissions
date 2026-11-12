func reverseWords(s string) string {
	if len(s) <= 1 {
		return s
	}
	words := strings.Split(s, " ")
	var reversed string

	for k, word := range words {
		chars := []rune(word)
		for l, r := 0, len(chars)-1; l < r; l, r = l+1, r-1 {
			chars[l], chars[r] = chars[r], chars[l]
		}
		if k > 0 {
			reversed += " "
		}
		reversed += string(chars)
	}
	return reversed
}
