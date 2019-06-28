func diStringMatch(S string) []int {
	min, max := 0, len(S)
	var ans []int
	for _, c := range S {
		if c == 'D' {
			ans = append(ans, max)
			max--
		} else if c == 'I' {
			ans = append(ans, min)
			min++
		}
	}
	ans = append(ans, min)
	return ans
}

