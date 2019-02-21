func checkAllRotten(grid [][]int) bool {
	for i := range grid {
		for _, x := range grid[i] {
			if x == 1 {
				return false
			}
		}
	}
	return true
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func orangesRotting(grid [][]int) int {
	if checkAllRotten(grid) {
		return 0
	}
	n := len(grid)
	m := len(grid[0])
	time := 0
	for {
		time++
		var found bool
		flag := make([][]bool, n)
		for i := 0; i < n; i++ {
			flag[i] = make([]bool, m)
		}

		for i := range grid {
			for j, x := range grid[i] {
				if x != 2 || flag[i][j] {
					continue
				}

				for _, direction := range directions {
					ni, nj := i+direction[0], j+direction[1]
					if ni < 0 || ni >= n || nj < 0 || nj >= m {
						continue
					}

					if grid[ni][nj] == 1 {
						grid[ni][nj] = 2
						flag[ni][nj] = true
						found = true
					}
				}
			}
		}

		if checkAllRotten(grid) {
			break
		}

		if !found {
			return -1
		}
	}
	return time
}

