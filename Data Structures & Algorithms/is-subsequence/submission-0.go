func isSubsequence(s string, t string) bool {
    sRunes := []rune(s)
    tRunes := []rune(t)
    tPointer := 0

    for _, sRune := range sRunes {
        found := false
        for j := tPointer; j < len(tRunes); j++ {
            tRune := tRunes[j]
            if tRune == sRune {
                tPointer = j+1
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    return true
}
