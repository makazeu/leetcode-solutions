func findJudge(N int, trust [][]int) int {
	trusted := make([]int, N+1)
	trusting := make([]bool, N+1)

	for _, t := range trust {
		trusting[t[0]] = true
		trusted[t[1]] ++
	}

	for i := 1; i <= N; i++ {
		if !trusting[i] && trusted[i] == N-1 {
			return i
		}
	}
	return -1
}
