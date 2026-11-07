func clearDigits(s string) string {
	bytes := []byte(s)
	var lettersToRemove int

	for i := len(bytes)-1; i >= 0; i-- {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			if lettersToRemove > 0 {
				bytes = append(bytes[:i], bytes[i+1:]...)
				lettersToRemove--
			}
		} else {
			bytes = append(bytes[:i], bytes[i+1:]...)
			lettersToRemove++
		}
	}
	return string(bytes)
}
