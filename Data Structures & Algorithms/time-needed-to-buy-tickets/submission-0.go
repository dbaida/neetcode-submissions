func timeRequiredToBuy(tickets []int, k int) int {
	var timeTaken int
    for {
		tickets[0]--
		timeTaken++
		if tickets[0] == 0 {
			if k == 0 {
				break
			} else {
				k--
			}
			tickets = tickets[1:len(tickets)]
		} else {
			if k == 0 {
				k = len(tickets)-1
			} else {
				k--
			}
			tickets = append(tickets[1:], tickets[0])
		}
	}
	return timeTaken
}