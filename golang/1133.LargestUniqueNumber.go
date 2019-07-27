func largestUniqueNumber(A []int) int {
	dict := make(map[int]int)
	for _, e := range A {
		dict[e]++
	}
	ans := -1
	for k, v := range dict {
		if v == 1 && k > ans {
			ans = k
		}
	}
	return ans
}

