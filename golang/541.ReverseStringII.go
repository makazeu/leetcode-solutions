func reverseStr(s string, k int) string {
	chars := strings.Split(s, "")
	l := len(chars)
	for i := 0; i < l; {
		stop := min(i+k, l)
		for j := i; j < (i+stop)/2; j++ {
			chars[j], chars[stop-j-1+i] = chars[stop-j-1+i], chars[j]
		}
		i += 2 * k
		if i >= l {
			break
		}
	}
	ans := ""
	for _, e := range chars {
		ans += e
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

