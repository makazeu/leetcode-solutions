func reverseParentheses(s string) string {
	s = reverse(s, 0, len(s)-1)

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' || s[i] == ')' {
			continue
		}
		res += string(s[i])
	}
	return res
}

func reverse(s string, start, end int) string {
	b := 0
	p := 0
	for i := start; i <= end; i++ {
		c := s[i]
		if c == '(' {
			if b == 0 {
				p = i
			}
			b--
		} else if c == ')' {
			b++
			if b == 0 {
				s = reverse(s, p+1, i-1)
			}
		}
	}

	res := ""
	for i := 0; i < start; i++ {
		res += string(s[i])
	}
	for i := end; i >= start; i-- {
		res += string(s[i])
	}
	for i := end + 1; i < len(s); i++ {
		res += string(s[i])
	}

	return res
}

