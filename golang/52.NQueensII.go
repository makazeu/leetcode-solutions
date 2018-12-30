func totalNQueens(n int) int {
	num := 0
	ans := make([]int, n+1)

	row := 1
	for {
		if row == n+1 {
			num++
			row--
			continue
		}
		if row == 0 {
			break
		}

		found := false
		for i := ans[row] + 1; i <= n; i++ {
			if !checkAvailability(row, i, ans) {
				continue
			}

			ans[row] = i
			found = true
			break
		}

		if found {
			row++
			if row <= n {
				ans[row] = 0
			}
		} else {
			row--
		}
	}

	return num
}

func checkAvailability(row, col int, ans []int) bool {
	for i := 1; i <= row-1; i++ {
		if ans[i] == col {
			return false
		}
		if row+col == i+ans[i] || row-col == i-ans[i] {
			return false
		}
	}
	return true
}
