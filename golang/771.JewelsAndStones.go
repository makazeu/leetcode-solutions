func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	for _, c := range J {
		m[c] = true
	}
	
	ans := 0
	for _, c := range S {
		if m[c] {
			ans ++
		}
	}
	return ans
}

