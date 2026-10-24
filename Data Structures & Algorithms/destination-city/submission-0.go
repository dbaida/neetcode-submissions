func destCity(paths [][]string) string {
	tripMap := make(map[string]string, len(paths))
	for _, path := range paths {
		tripMap[path[0]] = path[1]
	}

	dest := paths[0][1]
	for {
		nextDest := tripMap[dest]
		if nextDest == "" {
			return dest
		}
		dest = nextDest
	}
}
