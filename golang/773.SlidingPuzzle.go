const DigitNum = 6

func getDigitFromHigh(num, d int) int {
	div := int(math.Pow(10, float64(DigitNum-d)))
	mod := int(math.Pow(10, float64(DigitNum-d+1)))
	return (num % mod) / div
}

func swapDigit(num, i, j int) int {
	a := getDigitFromHigh(num, i)
	b := getDigitFromHigh(num, j)
	x := int(math.Pow(10.0, float64(DigitNum-i)))
	y := int(math.Pow(10.0, float64(DigitNum-j)))
	num += b*x - a*x
	num += a*y - b*y
	return num
}

func findZero(num int) int {
	for i := 1; i <= DigitNum; i++ {
		if getDigitFromHigh(num, i) == 0 {
			return i
		}
	}
	return 0
}

var Directions = []int{-1, -3, 1, 3}

func slidingPuzzle(board [][]int) int {
	initial := board[0][0]*100000 + board[0][1]*10000 + board[0][2]*1000
	initial += board[1][0]*100 + board[1][1]*10 + board[1][2]
	if initial == 123450 {
		return 0
	}

	used := make(map[int]bool)
	dist := make(map[int]int)
	dist[initial] = 0
	used[initial] = true

	queue := list.New()
	queue.PushBack(initial)

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		now := e.Value.(int)
		used[now] = false

		zero := findZero(now)
		for _, direction := range Directions {
			if direction == -1 && zero%3 == 1 {
				continue
			}
			if direction == 1 && zero%3 == 0 {
				continue
			}
			if direction == -3 && zero <= 3 {
				continue
			}
			if direction == 3 && zero >= 4 {
				continue
			}
			nxtZero := zero + direction
			next := swapDigit(now, zero, nxtZero)
			if next == 123450 {
				return dist[now] + 1
			}

			d, exists := dist[next]
			if exists && d <= dist[now]+1 {
				continue
			}
			dist[next] = dist[now] + 1

			if !used[next] {
				used[next] = true
				queue.PushBack(next)
			}
		}

	}

	return -1
}

