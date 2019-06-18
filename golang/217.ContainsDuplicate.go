func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, x := range nums {
		if m[x] {
			return true
		}
		m[x] = true
	}
	return false
}

