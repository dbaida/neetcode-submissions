func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToTmap := make(map[byte]byte)
	tToSmap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		char := s[i]
		if correspondingChar, ok := sToTmap[char]; !ok {
			if _, exists := tToSmap[t[i]]; exists {
				return false
			}

			sToTmap[char] = t[i]
			tToSmap[t[i]] = char
		} else {
			if correspondingChar != t[i] {
				return false
			}
		}
	}
	return true
}
