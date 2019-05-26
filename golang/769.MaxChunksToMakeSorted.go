func maxChunksToSorted(arr []int) int {
	ans := 0
	max := -1
	for i, e := range arr {
		if e > max {
			max = e
		}
		if max == i {
			ans++
		}
	}
	return ans
}

