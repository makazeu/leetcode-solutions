func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	for i := 1; i <= k; i++ {
		shift(&grid, n, m)
	}
	return grid
}

func shift(grid *[][]int, n, m int) {
	temp := make([][]int, n)
	for i, row := range *grid {
		temp[i] = make([]int, m)
		for j := 0; j < m-1; j++ {
			temp[i][j+1] = row[j]
		}
	}
	for i := 0; i < n-1; i++ {
		temp[i+1][0] = (*grid)[i][m-1]
	}
	temp[0][0] = (*grid)[n-1][m-1]
	*grid = temp
}

