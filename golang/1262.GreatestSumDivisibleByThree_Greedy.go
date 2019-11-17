func maxSumDivThree(nums []int) int {
	left := make([][]int, 3)
	for _, e := range nums {
		left[e%3] = append(left[e%3], e)
	}

	ans := 0
	// mod 3 == 0
	for _, e := range left[0] {
		ans += e
	}

	max := cal1(left)
	tmp := cal2(left)
	if tmp > max {
		max = tmp
	}
	return ans + max
}

func cal1(original [][]int) int {
	left := make([][]int, 3)
	for i := range original {
		left[i] = make([]int, len(original[i]))
		for j := range original[i] {
			left[i][j] = original[i][j]
		}
	}

	ans := 0
	sort.Ints(left[2])
	p2 := len(left[2]) - 1
	for p2 >= 1 {
		left[1] = append(left[1], left[2][p2]+left[2][p2-1])
		p2 -= 2
	}
	sort.Ints(left[1])
	p1 := len(left[1]) - 1

	for p1 >= 3 {
		ans += left[1][p1] + left[1][p1-1] + left[1][p1-2]
		p1 -= 3
	}

	max := 0
	if p1 == 2 {
		tmp := left[1][0] + left[1][1] + left[1][2]
		if tmp > max {
			max = tmp
		}
	}
	if p2 == 0 && p1 >= 0 {
		tmp := left[2][0] + left[1][p1]
		if tmp > max {
			max = tmp
		}
	}

	return ans + max
}

func cal2(original [][]int) int {
	left := make([][]int, 3)
	for i := range original {
		left[i] = make([]int, len(original[i]))
		for j := range original[i] {
			left[i][j] = original[i][j]
		}
	}

	ans := 0
	sort.Ints(left[1])
	p1 := len(left[1]) - 1
	for p1 >= 1 {
		left[2] = append(left[2], left[1][p1]+left[1][p1-1])
		p1 -= 2
	}
	sort.Ints(left[2])
	p2 := len(left[2]) - 1

	for p2 >= 3 {
		ans += left[2][p2] + left[2][p2-1] + left[2][p2-2]
		p2 -= 3
	}

	max := 0
	if p2 == 2 {
		tmp := left[2][0] + left[2][1] + left[2][2]
		if tmp > max {
			max = tmp
		}
	}
	if p1 == 0 && p2 >= 0 {
		tmp := left[1][0] + left[2][p2]
		if tmp > max {
			max = tmp
		}
	}

	return ans + max
}

