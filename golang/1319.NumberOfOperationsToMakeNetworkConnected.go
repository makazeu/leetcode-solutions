func makeConnected(n int, connections [][]int) int {
	ufs := make([]int, n)
	for i := 0; i < n; i++ {
		ufs[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if ufs[x] != x {
			ufs[x] = find(ufs[x])
		}
		return ufs[x]
	}
	merge := func(x int, y int) {
		x = find(x)
		y = find(y)
		ufs[y] = x
	}

	for _, con := range connections {
		merge(con[0], con[1])
	}

	hash := make(map[int]int)
	for i := 0; i < n; i++ {
		hash[find(i)]++
	}

	extra := len(hash) - 1
	used := 0
	for _, v := range hash {
		used += v - 1
	}
	if used+extra > len(connections) {
		return -1
	}
	return extra
}

