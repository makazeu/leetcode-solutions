func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	maxDifficulty := make([][]int, n+1)
	maxDifficulty[0] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		maxDifficulty[i] = make([]int, n+1)
		for j := i; j <= n; j++ {
			if j > i {
				maxDifficulty[i][j] = maxDifficulty[i][j-1]
			}
			if jobDifficulty[j-1] > maxDifficulty[i][j] {
				maxDifficulty[i][j] = jobDifficulty[j-1]
			}
		}
		//fmt.Println(maxDifficulty[i])
	}

	dp := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= d; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for i := 1; i <= d; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= j; k++ {
				if dp[i][j] > dp[i-1][j-k]+maxDifficulty[j-k+1][j] {
					dp[i][j] = dp[i-1][j-k] + maxDifficulty[j-k+1][j]
				}
			}
		}
		fmt.Println(dp[i])
	}
	return dp[d][n]
}

