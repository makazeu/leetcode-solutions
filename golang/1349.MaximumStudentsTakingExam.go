func maxStudents(seats [][]byte) int {
	m := len(seats)
	n := len(seats[0])
	maxNum := 1 << n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, maxNum)
	}

	masked := make([]int, m)
	for i := range masked {
		for j := 0; j < n; j++ {
			if seats[i][j] == '#' {
				masked[i] |= 1 << j
			}
		}
	}

	count1 := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		t := i
		c := 0
		for ; t > 0; t >>= 1 {
			c += t & 1
		}
		count1[i] = c
	}

	for k := range dp {
		for i := 0; i < maxNum; i++ {
			if i&masked[k] > 0 {
				continue
			}
			if i&(i<<1) > 0 {
				continue
			}
			max := 0
			if k >= 1 {
				for j := 0; j < maxNum; j++ {
					if (i<<1)&j > 0 {
						continue
					}
					if (i>>1)&j > 0 {
						continue
					}
					if max < dp[k-1][j] {
						max = dp[k-1][j]
					}
				}
			}
			dp[k][i] = max + count1[i]
		}
	}

	ans := 0
	for i := 0; i < maxNum; i++ {
		if ans < dp[m-1][i] {
			ans = dp[m-1][i]
		}
	}

	return ans
}

