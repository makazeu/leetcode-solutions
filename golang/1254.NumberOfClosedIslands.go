var directions = [][]int{
	{1, 0}, {-1, 0},
	{0, 1}, {0, -1},
}

func closedIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	flag := make([][]bool, n)
	for i := range flag {
		flag[i] = make([]bool, m)
	}
	ans := 0
	for i := range grid {
		for j := 0; j < m; j++ {
			if flag[i][j] {
				continue
			}
			if grid[i][j] == 0 && !dfs(grid, flag, n, m, i, j) {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, flag [][]bool, n, m, x, y int) bool {
	flag[x][y] = true

	var nx, ny int
	var res = false
	for _, direction := range directions {
		nx = x + direction[0]
		ny = y + direction[1]
		if nx == n || ny == m || nx == -1 || ny == -1 {
			res = true
			continue
		}
		if flag[nx][ny] {
			continue
		}
		if grid[nx][ny] == 0 && dfs(grid, flag, n, m, nx, ny) {
			res = true
		}
	}
	return res
}

