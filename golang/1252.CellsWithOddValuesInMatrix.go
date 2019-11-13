func oddCells(n int, m int, indices [][]int) int {
	row := make([]int, n)
	col := make([]int, m)
	for _, index := range indices {
		row[index[0]]++
		col[index[1]]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (row[i]+col[j])&1 == 1 {
				ans++
			}
		}
	}
	return ans
}

