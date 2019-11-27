func countServers(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	row := make([]int, n)
	col := make([]int, m)
	for i, r := range grid {
		for j, v := range r {
			if v == 0 {
				continue
			}
			row[i]++
			col[j]++
		}
	}

	ans := 0
	for i, r := range grid {
		for j, v := range r {
			if v == 0 {
				continue
			}
			if row[i] > 1 || col[j] > 1 {
				ans++
			}
		}
	}
	return ans
}

