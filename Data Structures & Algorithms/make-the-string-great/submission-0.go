func makeGood(s string) string {
	lowerToUpper := make(map[byte]byte)
	for char := byte('a'); char <= 'z'; char++ {
		lowerToUpper[char] = char - 'a' + 'A'
	}

	bytes := []byte(s)
	for i := 1; i < len(bytes); i++ {
		first := bytes[i-1]
		second := bytes[i]

		if isPair(lowerToUpper, first, second) || isPair(lowerToUpper, second, first) {
			bytes = append(bytes[:i-1], bytes[i+1:]...)
			i = 0
		}
	}
	return string(bytes)
}

func isPair(lowerToUpper map[byte]byte, first byte, second byte) bool {
	upper, ok := lowerToUpper[first]

	return ok && upper == second
}