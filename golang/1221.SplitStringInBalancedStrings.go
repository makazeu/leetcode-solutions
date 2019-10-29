func balancedStringSplit(s string) int {
	ans := 0
	x := 0
	for _, c := range s {
		if c == 'L' {
			x--
		} else if c == 'R' {
			x++
		}
		if x == 0 {
			ans++
		}
	}
	return ans
}

