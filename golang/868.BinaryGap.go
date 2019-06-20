func binaryGap(N int) int {
	last := -1
	ans := 0
	i := 0
	for N > 0 {
		i++
		m := N % 2
		N /= 2
		if m == 0 {
			continue
		}
		if last != -1 && i-last > ans {
			ans = i - last
		}
		last = i
	}
	return ans
}

