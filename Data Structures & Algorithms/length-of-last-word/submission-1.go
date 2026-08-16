func lengthOfLastWord(s string) int {
    const space = byte(' ')
    length := 0
    caughtLetter := false

    for i := len(s) - 1; i >= 0; i-- {
        char := s[i]
        if char != space {
            length++
            caughtLetter = true
        } else if caughtLetter {
            return length
        }
    }
	return length
}
