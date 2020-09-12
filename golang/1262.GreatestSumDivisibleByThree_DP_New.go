func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
		dp[i][1] = math.MinInt32
		dp[i][2] = math.MinInt32
	}
	for i, num := range nums {
		for j := 0; j < 3; j++ {
			if dp[i][j] > dp[i+1][j] {
				dp[i+1][j] = dp[i][j]
			}
			if dp[i][j]+num > dp[i+1][(j+num)%3] {
				dp[i+1][(j+num)%3] = dp[i][j] + num
			}
		}
	}
	return dp[n][0]
}

