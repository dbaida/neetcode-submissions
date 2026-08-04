func isAnagram(s string, t string) bool {
	size := len(s)
	if size != len(t) {
		return false
	}

	hashMapS := make(map[byte]int)
	hashMapT := make(map[byte]int)

	for i := 0; i < size; i++ {
		incrementCharCounter(hashMapS, s[i])
		incrementCharCounter(hashMapT, t[i])
	}

	if len(hashMapS) != len(hashMapT) {
		return false
	}

	for char, countS := range hashMapS {
		if countT, ok := hashMapT[char]; !ok || countS != countT {
			return false
		}
	}

	return true
}

func incrementCharCounter(hashMap map[byte]int, char byte) {
	if _, ok := hashMap[char]; !ok {
		hashMap[char] = 1
	}
	hashMap[char]++
}
