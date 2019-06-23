func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dist[0][0] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
			if i > 0 && dist[i-1][j] < dist[i][j] {
				dist[i][j] = dist[i-1][j]
			}
			if j > 0 && dist[i][j-1] < dist[i][j] {
				dist[i][j] = dist[i][j-1]
			}
			dist[i][j] += grid[i][j]
		}
	}
	return dist[m-1][n-1]
}

