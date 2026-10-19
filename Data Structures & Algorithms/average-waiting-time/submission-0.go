func averageWaitingTime(customers [][]int) float64 {
    var waitingTime int
	var chefReadyAt int

	for _, customer := range customers {
		arrival, timeToPrepare := customer[0], customer[1]

		if chefReadyAt < arrival {
			chefReadyAt = arrival
		}
		waitingTime += chefReadyAt + timeToPrepare - arrival
		chefReadyAt += timeToPrepare
	}

	return float64(waitingTime) / float64(len(customers))
}