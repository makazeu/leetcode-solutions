func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, num := range nums1 {
		m1[num]++
	}
	for _, num := range nums2 {
		m2[num]++
	}
	var ans []int
	for k, v := range m1 {
		for i := 1; i <= min(v, m2[k]); i++ {
			ans = append(ans, k)
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

