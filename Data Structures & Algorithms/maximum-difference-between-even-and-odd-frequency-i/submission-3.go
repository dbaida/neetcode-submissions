func maxDifference(s string) int {
	hashMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}

	maxOdd, minEven := 0, len(s)
	for _, num := range hashMap {
		if num % 2 == 0 {
			if num < minEven {
				minEven = num
			}
			continue
		}
		if num > maxOdd {
			maxOdd = num
		}
	}

	return maxOdd - minEven
}
