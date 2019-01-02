func uniquePaths(m int, n int) int {
	var matrix [2][]int
	matrix[0] = make([]int, n+1)
	matrix[1] = make([]int, n+1)

	matrix[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			matrix[i%2][j] = matrix[(i-1)%2][j] + matrix[i%2][j-1]
		}
	}
	return matrix[m%2][n]
}
