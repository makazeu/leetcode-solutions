func findLHS(nums []int) int {
	ans := 0
	m := make(map[int]int)
	for _, e := range nums {
		m[e]++
	}
	for k, v := range m {
		v2 := m[k+1]
		if v2 > 0 && v+v2 > ans {
			ans = v + v2
		}
	}
	return ans
}

