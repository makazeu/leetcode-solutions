func maximalSquare(matrix [][]byte) int {
	mat := make([][]int, len(matrix))
	for i := range mat {
		mat[i] = make([]int, len(matrix[i]))
	}

	max := 0
	for i := range matrix {
		for j, v := range matrix[i] {
			if v == '0' {
				mat[i][j] = 0
				continue
			}
			if i-1 >= 0 && j-1 >= 0 {
				mat[i][j] = minInt(mat[i-1][j], mat[i][j-1], mat[i-1][j-1]) + 1
			} else {
				mat[i][j] = 1
			}

			max = maxInt(max, mat[i][j]*mat[i][j])
		}
	}

	return max
}

func minInt(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
