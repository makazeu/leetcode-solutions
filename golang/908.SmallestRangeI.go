func smallestRangeI(A []int, K int) int {
	if len(A) <= 1 {
		return 0
	}
	max, min := A[0], A[0]
	for _, e := range A {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	ans := max - min - 2*K
	if ans < 0 {
		ans = 0
	}
	return ans
}

