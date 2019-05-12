func maxSumAfterPartitioning(A []int, K int) int {
	dp := make([]int, len(A)+1)
	for i := 1; i <= len(A); i++ {
		dp[i] = dp[i-1] + A[i-1]

		max := A[i-1]
		for j := 1; j <= K && i-j >= 0; j++ {
			if A[i-j] > max {
				max = A[i-j]
			}
			if dp[i-j]+max*j > dp[i] {
				dp[i] = dp[i-j] + max*j
			}
		}
	}
	return dp[len(A)]
}

