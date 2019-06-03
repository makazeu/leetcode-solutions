func partitionLabels(S string) []int {
	count := make([]int, 26)
	for _, c := range S {
		count[c-'a']++
	}

	var ans []int
	m := make(map[int]int)
	l := 0
	for _, c := range S {
		l++
		v := int(c - 'a')
		m[v]++
		if m[v] == count[v] {
			delete(m, v)
		}

		if len(m) == 0 {
			ans = append(ans, l)
			l = 0
		}
	}
	return ans
}

