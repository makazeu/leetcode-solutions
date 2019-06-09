type Pair struct{ x, y int }

func orderOfLargestPlusSign(N int, mines [][]int) int {
	m := make(map[Pair]bool)
	for _, e := range mines {
		m[Pair{e[0], e[1]}] = true
	}
	up := make([][]int, N)
	down := make([][]int, N)
	left := make([][]int, N)
	right := make([][]int, N)

	for i := 0; i < N; i++ {
		up[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if _, exists := m[Pair{i, j}]; exists {
				continue
			}
			up[i][j] = 1
			if i > 0 {
				up[i][j] += up[i-1][j]
			}
		}
	}

	for i := N - 1; i >= 0; i-- {
		down[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if _, exists := m[Pair{i, j}]; exists {
				continue
			}
			down[i][j] = 1
			if i < N-1 {
				down[i][j] += down[i+1][j]
			}
		}
	}

	for i := 0; i < N; i++ {
		left[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if _, exists := m[Pair{i, j}]; exists {
				continue
			}
			left[i][j] = 1
			if j > 0 {
				left[i][j] += left[i][j-1]
			}
		}
	}

	for i := 0; i < N; i++ {
		right[i] = make([]int, N)
		for j := N - 1; j >= 0; j-- {
			if _, exists := m[Pair{i, j}]; exists {
				continue
			}
			right[i][j] = 1
			if j < N-1 {
				right[i][j] += right[i][j+1]
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			n := min4(up[i][j], down[i][j], left[i][j], right[i][j])
			if n > ans {
				ans = n
			}
		}
	}
	return ans
}

func min4(a, b, c, d int) int {
	return min(min(a, b), min(c, d))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

