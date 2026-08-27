func isPathCrossing(path string) bool {
	loc := func(x, y int) string {
		return fmt.Sprintf("%d|%d", x, y)
	}
	x, y := 0, 0
    visits := make(map[string]bool)
	visits[loc(x, y)] = true
	
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		position := loc(x, y)
		if visits[position] {
			return true
		}
		visits[position] = true
	}
	return false
}
