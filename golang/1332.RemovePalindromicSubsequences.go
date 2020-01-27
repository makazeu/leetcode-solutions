func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return 2
		}
	}
	return 1
}

