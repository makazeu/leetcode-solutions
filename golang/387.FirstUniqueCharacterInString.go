func firstUniqChar(s string) int {
	m := make([]int, 26)
	pos := make([]int, 26)
	for i, c := range s {
		m[c-'a']++
		if m[c-'a'] == 1 {
			pos[c-'a'] = i
		}
	}

	p := len(s)
	for i, c := range m {
		if c == 1 && pos[i] < p {
			p = pos[i]
		}
	}

	if p == len(s) {
		return -1
	}
	return p
}

