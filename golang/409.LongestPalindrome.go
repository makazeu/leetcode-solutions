func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	
	ans := 0
	odd := 0
	for _, v := range m {
		if v % 2 == 1 {
			odd = 1
		}
		ans += v / 2 * 2
	}
	return ans + odd
}

