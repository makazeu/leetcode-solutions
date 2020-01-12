const MaxKey = 26

func minimumDistance(word string) int {
	dp := make([][][]int, len(word))
	for k := range dp {
		dp[k] = make([][]int, MaxKey)
		for i := range dp[k] {
			dp[k][i] = make([]int, MaxKey)
			for j := range dp[k][i] {
				dp[k][i][j] = len(word) * MaxKey
			}
		}
	}
	for i := 0; i < MaxKey; i++ {
		dp[0][word[0]-'A'][i] = 0
		dp[0][i][word[0]-'A'] = 0
	}

	for k := 1; k < len(word); k++ {
		m := int(word[k] - 'A')
		// Uses finger 1
		for j := 0; j < MaxKey; j++ {
			for i := 0; i < MaxKey; i++ {
				if dp[k][m][j] <= dp[k-1][i][j] {
					continue
				}
				dp[k][m][j] = Min(dp[k][m][j], dp[k-1][i][j]+getDistance(m, i))
			}
		}

		// Uses finger 2
		for i := 0; i < MaxKey; i++ {
			for j := 0; j < MaxKey; j++ {
				if dp[k][i][m] <= dp[k-1][i][j] {
					continue
				}
				dp[k][i][m] = Min(dp[k][i][m], dp[k-1][i][j]+getDistance(m, j))
			}
		}
	}

	ans := math.MaxInt32
	for i := 0; i < MaxKey; i++ {
		for j := 0; j < MaxKey; j++ {
			ans = Min(ans, dp[len(word)-1][i][j])
		}
	}
	return ans
}

func getDistance(a, b int) int {
	ax, ay := a/6, a%6
	bx, by := b/6, b%6
	return Abs(ax-bx) + Abs(ay-by)
}

func Min(x, y int) int {
	return If(x <= y, x, y).(int)
}

func Abs(x int) int {
	return If(x >= 0, x, -x).(int)
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

