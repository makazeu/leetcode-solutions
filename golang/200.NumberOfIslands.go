var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = 0

	for i := 0; i < 4; i++ {
		tx, ty := x+direction[i][0], y+direction[i][1]
		if tx < 0 || tx >= len(grid) || ty < 0 || ty >= len(grid[tx]) {
			continue
		}
		if grid[tx][ty] == '1' {
			dfs(grid, tx, ty)
		}
	}
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}