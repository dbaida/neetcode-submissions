func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)
    pieces := strings.Split(trimmed, " ")

    if len(pieces) == 0 {
        return 0
    }
    return len(pieces[len(pieces) - 1])
}
