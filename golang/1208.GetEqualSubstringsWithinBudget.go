func equalSubstring(s string, t string, maxCost int) int {
	a := make([]int, len(s))
	for i := range a {
		a[i] = abs(int(s[i]) - int(t[i]))
	}

	ans := 0
	left := 0
	cost := 0
	for i := range a {
		cost += a[i]
		for cost > maxCost {
			cost -= a[left]
			left++
		}
		if i-left+1 > ans {
			ans = i - left + 1
		}
	}

	return ans
}

