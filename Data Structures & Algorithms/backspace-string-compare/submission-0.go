func backspaceCompare(s string, t string) bool {
	return string(runSimulation(s)) == string(runSimulation(t))
}

func runSimulation(str string) []byte {
	editor := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == '#' {
			if len(editor) > 0 {
				editor = editor[:len(editor)-1]
			}
			continue
		}
		editor = append(editor, str[i])
	}
	return editor
}
