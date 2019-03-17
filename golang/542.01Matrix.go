var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Node struct {
	x, y int
}

func updateMatrix(matrix [][]int) [][]int {
	dist := make([][]int, len(matrix))
	for i, row := range matrix {
		dist[i] = make([]int, len(row))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	// init
	queue := list.New()
	for i, row := range matrix {
		for j := range row {
			if row[j] != 0 {
				continue
			}
			dist[i][j] = 0
			node := &Node{i, j}
			queue.PushBack(node)
		}
	}

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)

		node := e.Value.(*Node)
		d := dist[node.x][node.y] + 1
		for _, direction := range directions {
			x := node.x + direction[0]
			y := node.y + direction[1]

			if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) {
				continue
			}
			if dist[x][y] <= d {
				continue
			}
			dist[x][y] = d

			queue.PushBack(&Node{x, y})
		}
	}

	return dist
}

