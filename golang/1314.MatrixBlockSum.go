func matrixBlockSum(mat [][]int, K int) [][]int {
	R, C := len(mat), len(mat[0])
	ans := make([][]int, R)
	for i := range mat {
		for j := range mat[i] {
			if j > 0 {
				mat[i][j] += mat[i][j-1]
			}
		}
		ans[i] = make([]int, C)
	}

	for c := range mat[0] {
		var ca, cb int
		if ca = c - K; ca < 0 {
			ca = 0
		}
		if cb = c + K; cb >= C {
			cb = C - 1
		}
		for i := 0; i <= Min(R-1, K); i++ {
			ans[0][c] += Sum(mat, i, ca, cb)
		}
		for r := 1; r < R; r++ {
			ans[r][c] = ans[r-1][c]
			if r-K-1 >= 0 {
				ans[r][c] -= Sum(mat, r-K-1, ca, cb)
			}
			if r+K < R {
				ans[r][c] += Sum(mat, r+K, ca, cb)
			}
		}
	}
	return ans
}

func Sum(mat [][]int, r, c1, c2 int) int {
	ans := mat[r][c2]
	if c1-1 >= 0 {
		ans -= mat[r][c1-1]
	}
	return ans
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

