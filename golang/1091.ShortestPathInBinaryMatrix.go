var directions = [][]int{
	{-1, 0}, {-1, 1}, {-1, -1},
	{0, 1}, {0, -1},
	{1, 1}, {1, 0}, {1, -1},
}

type Point struct {
	x, y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	N := len(grid)
	if grid[0][0] == 1 || grid[N-1][N-1] == 1 {
		return -1
	}
	dist := make([][]int, N)
	flag := make([][]bool, N)
	for i := range dist {
		dist[i] = make([]int, N)
		flag[i] = make([]bool, N)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 1
	flag[0][0] = true

	queue := list.New()
	queue.PushBack(Point{0, 0})
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		point := e.Value.(Point)
		flag[point.x][point.y] = false

		var x, y, d int
		for _, direction := range directions {
			x = point.x + direction[0]
			y = point.y + direction[1]
			if x < 0 || y < 0 || x == N || y == N {
				continue
			}
			if grid[x][y] == 1 {
				continue
			}
			d = dist[point.x][point.y] + 1
			if d >= dist[x][y] {
				continue
			}
			dist[x][y] = d
			if !flag[x][y] {
				queue.PushBack(Point{x, y})
				flag[x][y] = true
			}
		}
	}

	if dist[N-1][N-1] != math.MaxInt32 {
		return dist[N-1][N-1]
	} else {
		return -1
	}
}

