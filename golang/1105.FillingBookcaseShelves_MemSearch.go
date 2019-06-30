func minHeightShelves(books [][]int, shelf_width int) int {
	N := len(books)
	thinSum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		thinSum[i] = thinSum[i-1] + books[i-1][0]
	}

	result := make([][]int, N)
	for i := range books {
		result[i] = make([]int, N)
	}
	dfs(0, N-1, thinSum, result, books, shelf_width)
	return result[0][N-1]
}

func dfs(start, stop int, thinSum []int, result, books [][]int, width int) {
	if result[start][stop] != 0 {
		return
	}
	thin := thinSum[stop+1] - thinSum[start]
	if thin <= width {
		max := 0
		for i := start; i <= stop; i++ {
			if books[i][1] > max {
				max = books[i][1]
			}
		}
		result[start][stop] = max
		return
	}
	for i := start; i <= stop-1; i++ {
		dfs(start, i, thinSum, result, books, width)
		dfs(i+1, stop, thinSum, result, books, width)
		a := result[start][i]
		b := result[i+1][stop]

		if result[start][stop] == 0 || result[start][stop] > a+b {
			result[start][stop] = a + b
		}
	}
}

