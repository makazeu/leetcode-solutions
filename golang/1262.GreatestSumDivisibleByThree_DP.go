func maxSumDivThree(nums []int) int {
	left := make([][]int, 3)
	for _, e := range nums {
		left[e%3] = append(left[e%3], e)
	}

	ans := 0
	// mod 3 == 0
	for _, e := range left[0] {
		ans += e
	}

	sort.Ints(left[1])
	sort.Ints(left[2])
	p1, p2 := len(left[1]), len(left[2])
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = make([]int, p2+1)
	}

	var tmp int
	for i := 0; i <= p1; i++ {
		for j := 0; j <= p2; j++ {
			if i >= 3 {
				tmp = dp[(i-3)%4][j] + left[1][i-3] + left[1][i-2] + left[1][i-1]
				if tmp > dp[i%4][j] {
					dp[i%4][j] = tmp
				}
			}

			if j >= 3 {
				tmp = dp[i%4][j-3] + left[2][j-3] + left[2][j-2] + left[2][j-1]
				if tmp > dp[i%4][j] {
					dp[i%4][j] = tmp
				}
			}

			if i >= 1 && j >= 1 {
				tmp = dp[(i-1)%4][j-1] + left[1][i-1] + left[2][j-1]
				if tmp > dp[i%4][j] {
					dp[i%4][j] = tmp
				}
			}
		}
	}
	return ans + dp[p1%4][p2]
}

