func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, num := range nums1 {
		m1[num] = true
	}
	for _, num := range nums2 {
		m2[num] = true
	}
	var ans []int
	for k := range m1 {
		if m2[k] {
			ans = append(ans, k)
		}
	}
	return ans
}

