func minOperations(logs []string) int {
	var ops int
	isInMainFolder := true

	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" || (logs[i] == "../" && isInMainFolder) {
			continue
		}
		if logs[i] == "../" {
			ops--
			isInMainFolder = ops == 0
			continue
		}
		isInMainFolder = false
		ops++
	}
	return ops
}
