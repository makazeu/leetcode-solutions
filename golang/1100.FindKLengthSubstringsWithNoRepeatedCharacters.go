func numKLenSubstrNoRepeats(S string, K int) int {
	if len(S) < K {
		return 0
	}
	ans := 0
	dup := 0
	chars := make([]int, 26)
	for i := 0; i < K; i++ {
		c := S[i] - 'a'
		chars[c]++
		if chars[c] == 2 {
			dup++
		}
	}
	if dup == 0 {
		ans++
	}

	for i := 1; i <= len(S)-K; i++ {
		c := S[i-1] - 'a'
		chars[c]--
		if chars[c] == 1 {
			dup--
		}

		c = S[i+K-1] - 'a'
		chars[c] ++
		if chars[c] == 2 {
			dup++
		}
		if dup == 0 {
			ans++
		}
	}
	return ans
}

