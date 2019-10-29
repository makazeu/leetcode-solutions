type Position struct {
	x, y int
}

var directions = [][]int{
	{1, 0}, {1, -1}, {1, 1},
	{0, -1}, {0, 1},
	{-1, 0}, {-1, -1}, {-1, 1},
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var ans [][]int
	m := make(map[Position]bool)
	for _, queen := range queens {
		m[Position{queen[0], queen[1]}] = true
	}
	var x, y int
	for _, direction := range directions {
		for i := 1; ; i++ {
			x = king[0] + i*direction[0]
			y = king[1] + i*direction[1]

			if x < 0 || y < 0 || x >= 8 || y >= 8 {
				break
			}

			if m[Position{x, y}] {
				ans = append(ans, []int{x, y})
				break
			}
		}
	}
	return ans
}

