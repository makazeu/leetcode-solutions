func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 3)
		dp[i][1] = math.MinInt32
		dp[i][2] = math.MinInt32
	}
	for i, num := range nums {
		for j := 0; j < 3; j++ {
			if dp[i%2][j] > dp[(i+1)%2][j] {
				dp[(i+1)%2][j] = dp[i%2][j]
			}
			if dp[i%2][j]+num > dp[(i+1)%2][(j+num)%3] {
				dp[(i+1)%2][(j+num)%3] = dp[i%2][j] + num
			}
		}
	}
	return dp[n%2][0]
}

