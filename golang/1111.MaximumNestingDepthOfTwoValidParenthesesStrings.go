func maxDepthAfterSplit(seq string) []int {
	s := 0
	m := 0
	for _, c := range seq {
		if c == '(' {
			s++
			if s > m {
				m = s
			}
		} else {
			s--
		}
	}
	mid := m / 2
	ans := make([]int, len(seq))
	for i, c := range seq {
		if c == '(' {
			s++
			if s > mid {
				ans[i] = 1
			}
		} else {
			if s > mid {
				ans[i] = 1
			}
			s--
		}
	}
	return ans
}

