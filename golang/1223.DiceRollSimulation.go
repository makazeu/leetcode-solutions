const MOD = 1000000007

func dieSimulator(n int, rollMax []int) int {
	dp := make([][][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([][]int, 6)
		for j := range dp[i] {
			dp[i][j] = make([]int, rollMax[j]+1)
		}
	}
	for i := range dp[1] {
		dp[1][i][1] = 1
	}
	for i := 1; i < n; i++ {
		for j := range dp[i] {
			for k := 1; k <= rollMax[j]; k++ {
				if dp[i][j][k] == 0 {
					continue
				}
				for x := 0; x < 6; x++ {
					if j != x {
						dp[i+1][x][1] = (dp[i+1][x][1] + dp[i][j][k]) % MOD
					} else if k < rollMax[j] {
						dp[i+1][x][k+1] = (dp[i+1][x][k+1] + dp[i][j][k]) % MOD
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 6; i++ {
		for j := 1; j <= rollMax[i]; j++ {
			ans = (ans + dp[n][i][j]) % MOD
		}
	}
	return ans
}

