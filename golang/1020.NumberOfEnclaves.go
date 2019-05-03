type square struct {
	x, y int
}

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func numEnclaves(A [][]int) int {
	ans := 0
	visited := make([][]bool, len(A))
	for i := range A {
		visited[i] = make([]bool, len(A[i]))
	}

	for i := range A {
		for j := range A[i] {
			if A[i][j] == 0 || visited[i][j] {
				continue
			}
			num, flag := bfs(A, visited, i, j)
			if flag {
				ans += num
			}
		}
	}

	return ans
}

func bfs(A [][]int, V [][]bool, sx, sy int) (int, bool) {
	flag := true
	num := 1

	queue := list.New()
	queue.PushBack(square{sx, sy})
	V[sx][sy] = true

	for queue.Len() > 0 {
		sq := queue.Front()
		queue.Remove(sq)
		x, y := sq.Value.(square).x, sq.Value.(square).y

		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			if nx < 0 || nx >= len(A) || ny < 0 || ny >= len(A[x]) {
				flag = false
				continue
			}

			if A[nx][ny] == 1 && !V[nx][ny] {
				V[nx][ny] = true
				queue.PushBack(square{nx, ny})
				num++
			}
		}
	}

	return num, flag
}

