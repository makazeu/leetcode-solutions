func twoSumLessThanK(A []int, K int) int {
	ans := -1
	for i, x := range A {
		for j, y := range A {
			if i == j {
				continue
			}
			if x+y >= K {
				continue
			}
			if x+y > ans {
				ans = x + y
			}
		}
	}
	return ans
}

