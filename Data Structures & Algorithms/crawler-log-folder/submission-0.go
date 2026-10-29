func minOperations(logs []string) int {
	path := make([]string, 0)

	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" || (logs[i] == "../" && len(path) == 0) {
			continue
		}
		if logs[i] == "../" {
			path = path[:len(path)-1]
			continue
		}
		path = append(path, logs[i])
	}

	return len(path)
}
