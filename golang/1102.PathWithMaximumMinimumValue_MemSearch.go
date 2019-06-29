type Item struct {
	x, y int
}

var directions = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func maximumMinimumPath(A [][]int) int {
	R, C := len(A), len(A[0])
	dist := make([][]int64, R)
	queued := make([][]bool, R)
	for i := range A {
		dist[i] = make([]int64, C)
		queued[i] = make([]bool, C)
	}

	dist[0][0] = int64(A[0][0])
	queued[0][0] = true
	queue := list.New()
	queue.PushBack(&Item{0, 0})
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		item := e.Value.(*Item)
		x, y := item.x, item.y
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			if nx < 0 || nx >= R || ny < 0 || ny >= C {
				continue
			}
			d := min(dist[x][y], int64(A[nx][ny]))
			if d <= dist[nx][ny] {
				continue
			}
			dist[nx][ny] = d
			if queued[nx][ny] {
				continue
			}
			queue.PushBack(&Item{nx, ny})
		}
	}
	return int(dist[R-1][C-1])
}

func min(x, y int64) int64 {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

