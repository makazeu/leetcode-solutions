func longestWPI(hours []int) int {
	l := len(hours)
	sum := make([]int, l+1)
	for i, e := range hours {
		sum[i+1] = sum[i] + If(e > 8, 1, -1).(int)
	}

	h := make(map[int]int)
	for i := l; i >= 1; i-- {
		if h[sum[i]] == 0 {
			h[sum[i]] = i
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		target := sum[i] + 1
		now := 0
		if sum[l] >= target {
			now = l - i
		} else if h[target] > i {
			now = h[target] - i
		}
		if now > ans {
			ans = now
		}
	}

	return ans
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

