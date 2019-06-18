const (
	Buy = iota
	Sell
)

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices)+1)
	dp[0] = make([]int, 2)
	ans := 0
	for i := 1; i <= len(prices); i++ {
		dp[i] = make([]int, 2)
		for j := 0; j <= i-2; j++ {
			dp[i][Buy] = max(dp[i][Buy], dp[j][Sell])
		}
		for j := 1; j <= i-1; j++ {
			if prices[j-1] >= prices[i-1] {
				continue
			}
			dp[i][Sell] = max(dp[i][Sell], dp[j][Buy]+(prices[i-1]-prices[j-1]))
		}
		ans = max(ans, dp[i][Sell])
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

