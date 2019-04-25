func longestArithSeqLength(A []int) int {
	ans := 0
	dp := make([]map[int]int, len(A))
	for i := range A {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			dif := A[i] - A[j]
			v := dp[j][dif] + 1
			if v < 2 {
				v = 2
			}
			if v > dp[i][dif] {
				dp[i][dif] = v
			}

			if v > ans {
				ans = v
			}
		}
	}

	return ans
}

