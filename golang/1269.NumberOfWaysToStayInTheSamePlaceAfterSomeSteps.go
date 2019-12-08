func numWays(steps int, arrLen int) int {
	dp := make([][]int, 2)
	if arrLen > 2*steps+1 {
		arrLen = 2*steps + 1
	}
	dp[0] = make([]int, arrLen)
	dp[1] = make([]int, arrLen)

	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		x := i & 1
		y := (i - 1) & 1
		for j := 0; j < arrLen; j++ {
			dp[x][j] = dp[y][j]
			if j > 0 {
				dp[x][j] = (dp[x][j] + dp[y][j-1]) % 1000000007
			}
			if j < arrLen-1 {
				dp[x][j] = (dp[x][j] + dp[y][j+1]) % 1000000007
			}
		}
	}
	return dp[steps&1][0]
}

