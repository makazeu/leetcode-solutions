func smallestCommonElement(mat [][]int) int {
	n := len(mat)
	m := make(map[int]int)
	for _, row := range mat {
		for _, x := range row {
			m[x]++
		}
	}

	ans := math.MaxInt32
	for k, v := range m {
		if v == n && k < ans {
			ans = k
		}
	}

	if ans == math.MaxInt32 {
		return -1
	} 
	return ans
}

