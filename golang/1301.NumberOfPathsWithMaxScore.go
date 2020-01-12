const MOD = 1000000007

func pathsWithMaxScore(board []string) []int {
	N := len(board)
	maxDp := make([][]int, N)
	numDp := make([][]int, N)
	for i := 0; i < N; i++ {
		maxDp[i] = make([]int, N)
		numDp[i] = make([]int, N)
	}
	numDp[0][0] = 1

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if i > 0 && board[i-1][j] != 'X' && numDp[i-1][j] > 0 {
				maxDp[i][j] = max(maxDp[i][j], maxDp[i-1][j])
			}
			if j > 0 && board[i][j-1] != 'X' && numDp[i][j-1] > 0 {
				maxDp[i][j] = max(maxDp[i][j], maxDp[i][j-1])
			}
			if i > 0 && j > 0 && board[i-1][j-1] != 'X' && numDp[i-1][j-1] > 0 {
				maxDp[i][j] = max(maxDp[i][j], maxDp[i-1][j-1])
			}

			if i > 0 && board[i-1][j] != 'X' && maxDp[i][j] == maxDp[i-1][j] {
				numDp[i][j] = (numDp[i][j] + numDp[i-1][j]) % MOD
			}
			if j > 0 && board[i][j-1] != 'X' && maxDp[i][j] == maxDp[i][j-1] {
				numDp[i][j] = (numDp[i][j] + numDp[i][j-1]) % MOD
			}
			if i > 0 && j > 0 && board[i-1][j-1] != 'X' && maxDp[i][j] == maxDp[i-1][j-1] {
				numDp[i][j] = (numDp[i][j] + numDp[i-1][j-1]) % MOD
			}

			if !(i == 0 && j == 0) && !(i == N-1 && j == N-1) {
				maxDp[i][j] += int(board[i][j] - '0')
			}
		}
	}
	return []int{maxDp[N-1][N-1], numDp[N-1][N-1]}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

