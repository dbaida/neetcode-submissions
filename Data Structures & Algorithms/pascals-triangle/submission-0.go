func generate(numRows int) [][]int {
	var triangle [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
				continue
			}
			row = append(row, triangle[i-1][j-1] + triangle[i-1][j])
		}
		triangle = append(triangle, row)
	}
	return triangle
}
