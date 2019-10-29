func uniqueOccurrences(arr []int) bool {
	m1 := make(map[int]int)
	for _, x := range arr {
		m1[x]++
	}

	m := make(map[int]bool)
	for _, v := range m1 {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

