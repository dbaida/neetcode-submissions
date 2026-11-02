func getRow(rowIndex int) []int {
	rows := make([][]int, 0)
	rows = append(rows, []int{1})

    for i := 1; i <= rowIndex; i++ {
		rows = append(rows, []int{1})
		for j := 1; j < i; j++ {
			rows[i] = append(rows[i], rows[i-1][j-1] + rows[i-1][j])
		}
		rows[i] = append(rows[i], 1)
	}
	return rows[rowIndex]
}