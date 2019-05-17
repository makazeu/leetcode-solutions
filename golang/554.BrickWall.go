func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	ans := 0
	for _, w := range wall {
		sum := 0
		for i := 0; i < len(w)-1; i++ {
			sum += w[i]
			n := m[sum] + 1
			m[sum] = n

			if n > ans {
				ans = n
			}
		}
	}

	return len(wall) - ans
}

