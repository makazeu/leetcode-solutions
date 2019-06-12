func findAnagrams(s string, p string) []int {
	var ans []int
	if len(s) < len(p) {
		return ans
	}
	m := make(map[uint8]int)
	for i := range p {
		m[p[i]]--
	}
	for i := 0; i < len(p); i++ {
		m[s[i]]++
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
	}
	if len(m) == 0 {
		ans = append(ans, 0)
	}

	l := len(p)
	for i := 0; i < len(s)-l; i++ {
		m[s[i]]--
		m[s[i+l]]++
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
		if m[s[i+l]] == 0 {
			delete(m, s[i+l])
		}
		if len(m) == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

