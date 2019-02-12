func find(ufs []int, x int) int {
	if ufs[x] != x {
		ufs[x] = find(ufs, ufs[x])
		return ufs[x]
	} else {
		return x
	}
}

func union(ufs []int, x, y int) {
	x = find(ufs, x)
	y = find(ufs, y)
	ufs[y] = x
}

func equationsPossible(equations []string) bool {
	ufs := make([]int, 26)
	for i := 0; i <= 25; i++ {
		ufs[i] = i
	}

	for _, equation := range equations {
		if equation[1] != '=' {
			continue
		}
		x := int(equation[0] - 'a')
		y := int(equation[3] - 'a')
		union(ufs, x, y)
	}

	for _, equation := range equations {
		if equation[1] == '=' {
			continue
		}
		x := int(equation[0] - 'a')
		y := int(equation[3] - 'a')
		x = find(ufs, x)
		y = find(ufs, y)
		if x == y {
			return false
		}
	}
	return true
}
