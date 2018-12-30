func solveNQueens(n int) [][]string {
	var answers [][]string
	ans := make([]int, n+1)

	row := 1
	for {
		if row == n+1 {
			answers = append(answers, generateAnswer(n, ans))
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

	return answers
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

func generateAnswer(n int, ans []int) (canvas []string) {
	for i := 1; i <= n; i++ {
		str := strings.Repeat(".", ans[i]-1)
		str = str + "Q"
		str = str + strings.Repeat(".", n-ans[i])
		canvas = append(canvas, str)
	}
	return
}
