func isValid(s string) bool {
	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
    closeStack := make([]rune, 0)
	inputRunes := []rune(s)
    var nextClosingBracket rune
	for _, char := range inputRunes {
		expectedCloseBracket, isOpening := brackets[char]
		if isOpening {
            closeStack = append(closeStack, expectedCloseBracket)
			continue
		}
        if len(closeStack) == 0 {
            return false
        }
        // suppose it's closing bracket
        nextClosingBracket, closeStack = pop(closeStack)
        if char != nextClosingBracket {
            return false
        }
	}
	return len(closeStack) == 0
}

func pop(stack []rune) (rune, []rune) {
    idx := len(stack)-1
    value := stack[idx]
    stack = stack[:idx]

    return value, stack
}
