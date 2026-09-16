func largestGoodInteger(num string) string {
	var seqCounts [10]int
	largest := -1

	for i := 1; i < len(num); i++ {
		prev := num[i-1]
		curr := num[i]
		currInt, _ := strconv.Atoi(string(curr))

		if prev == curr {
			seqCounts[currInt]++
			if seqCounts[currInt] > 1 && currInt > largest {
				largest = currInt
			}
		} else {
			seqCounts[currInt] = 0
		}
	}

	if largest == -1 {
		return ""
	}

	return strings.Repeat(strconv.Itoa(largest), 3)
}
