func calPoints(operations []string) int {
	scores := make([]int, 0)
	var points int
	for _, op := range operations {
		if op == "+" {
			score := scores[len(scores)-2] + scores[len(scores)-1]
			scores = append(scores, score)
			points += score
			continue
		}
		if op == "C" || op == "D" {
			if op == "C" {
				points -= scores[len(scores)-1]
				scores = scores[:len(scores)-1]
				continue
			}
			score := scores[len(scores)-1] * 2
			scores = append(scores, score)
			points += score
			continue
		}

		score, _ := strconv.Atoi(op)
		scores = append(scores, score)
		points += score

	}
	return points
}
