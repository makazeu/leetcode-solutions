func isToeplitzMatrix(matrix [][]int) bool {
	N, M := len(matrix), len(matrix[0])
	for i := 0; i < N; i++ {
		v := matrix[i][0]
		for x, y := i, 0; x < N && y < M; x, y = x+1, y+1 {
			if matrix[x][y] != v {
				return false
			}
		}
	}
	for i := 1; i < M; i++ {
		v := matrix[0][i]
		for x, y := 0, i; x < N && y < M; x, y = x+1, y+1 {
			if matrix[x][y] != v {
				return false
			}
		}
	}
	return true
}

