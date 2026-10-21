func isValidSudoku(board [][]byte) bool {
	colsDigits := make(map[int]map[byte]bool)
	rowsDigits := make(map[int]map[byte]bool)

	// check cols & rows
	for i := 0; i < 9; i++ {
		if rowsDigits[i] == nil {
			rowsDigits[i] = make(map[byte]bool)
		}

		for j := 0; j < 9; j++ {
			if colsDigits[j] == nil {
				colsDigits[j] = make(map[byte]bool)
			}
			v := board[i][j]

			if isDigit(v) {
				if colsDigits[j][v] {
					return false
				}
				colsDigits[j][v] = true

				if rowsDigits[i][v] {
					return false
				}
				rowsDigits[i][v] = true
			}
		}
	}

	// check sub-boxes
	nums := make(map[byte]bool)
	for boxIdx := 0; boxIdx < 9; boxIdx++ {
		startRow := (boxIdx / 3) * 3
		startCol := (boxIdx % 3) * 3

		for r := startRow; r < startRow+3; r++ {
			for c := startCol; c < startCol+3; c++ {
				v := board[r][c]
				if !isDigit(v) {
					continue
				}
				if nums[v] {
					return false
				}
				nums[v] = true
			}
		}
		clear(nums)
	}

	return true
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
