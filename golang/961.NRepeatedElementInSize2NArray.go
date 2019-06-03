func repeatedNTimes(A []int) int {
	m := make(map[int]bool)
	for _, x := range A {
		if m[x] {
			return x
		}
		m[x] = true
	}
	return 0
}

