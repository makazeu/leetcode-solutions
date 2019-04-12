func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}

	ans := ""
	m := 0
	for {
		if N == 0 {
			break
		}
		if N % -2 != 0 {
			m = 1
			N = (N-1) / -2
		} else {
			m = 0
			N = N / -2
		}
		ans = strconv.Itoa(m) + ans
	}

	return ans
}

