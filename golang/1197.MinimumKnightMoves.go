var directions = [][]int{
	{-1, 2}, {1, 2}, {2, 1}, {2, -1},
	{1, -2}, {-1, -2}, {-2, -1}, {-2, 1},
}

type Item struct {
	x, y int
}

func minKnightMoves(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	m := make(map[int]map[int]int)
	flag := make(map[int]map[int]bool)
	m[0] = make(map[int]int)
	flag[0] = make(map[int]bool)
	m[0][0] = 0
	flag[0][0] = true

	q := list.New()
	q.PushBack(&Item{0, 0})

	for q.Len() > 0 {
		e := q.Front()
		q.Remove(e)
		item := e.Value.(*Item)
		cx, cy := item.x, item.y
		flag[cx][cy] = false
		cd := m[cx][cy]
		var nx, ny int
		for _, direction := range directions {
			nx = cx + direction[0]
			ny = cy + direction[1]

			if abs(nx) > 2*abs(x) && abs(ny) > 2*abs(y) {
				continue
			}
			if _, exists := m[nx]; !exists {
				m[nx] = make(map[int]int)
				flag[nx] = make(map[int]bool)
			}
			nd, exists := m[nx][ny]
			if !exists || nd > cd+1 {
				m[nx][ny] = cd + 1
				if nx == x && ny == y {
					return m[nx][ny]
				}
				if !flag[nx][ny] {
					q.PushBack(&Item{nx, ny})
					flag[nx][ny] = true
				}
			}
		}
	}
	panic("no answer exists")
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

