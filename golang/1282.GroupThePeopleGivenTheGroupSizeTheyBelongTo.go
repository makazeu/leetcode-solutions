func groupThePeople(groupSizes []int) [][]int {
	var group [][]int
	h := make(map[int]int)
	for i, k := range groupSizes {
		var x int
		x, exists := h[k]
		if !exists {
			group = append(group, make([]int, 0))
			x = len(group) - 1
			h[k] = x
		}
		group[x] = append(group[x], i)
		if len(group[x]) == k {
			delete(h, k)
		}
	}
	return group
}

