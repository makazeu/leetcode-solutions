var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func getNextDirection(cur, turn int) int {
	cur += turn
	if cur < 0 {
		cur = len(directions) - 1
	}
	return cur % len(directions)
}

func isRobotBounded(instructions string) bool {
	x, y, d := 0, 0, 0
	for i := 1; i <= 4; i++ {
		x, y, d = completeInstructions(x, y, d, &instructions)

		if x == 0 && y == 0 && d == 0 {
			return true
		}
	}

	return false
}

func completeInstructions(x, y, d int, instructions *string) (int, int, int) {
	for _, e := range *instructions {
		switch e {
		case 'G':
			x += directions[d][0]
			y += directions[d][1]
		case 'L':
			d = getNextDirection(d, -1)
		case 'R':
			d = getNextDirection(d, 1)
		}
	}
	return x, y, d
}

