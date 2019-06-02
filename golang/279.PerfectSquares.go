func numSquares(n int) int {
	var square []int
	for i := 1; i*i <= n; i++ {
		square = append(square, i*i)
	}
	best := make(map[int]int)
	flag := make(map[int]bool)

	queue := list.New()
	queue.PushBack(0)
	ans := n
	for queue.Len() > 0 {
		q := queue.Front()
		queue.Remove(q)
		x := q.Value.(int)
		flag[x] = false

		if best[x]+1 >= ans {
			continue
		}
		for _, k := range square {
			if x+k > n {
				continue
			}
			nxt := best[x] + 1
			if best[x+k] != 0 && best[x+k] <= nxt {
				continue
			}
			best[x+k] = nxt
			if x+k == n && nxt < ans {
				ans = nxt
			}

			if !flag[x+k] {
				queue.PushBack(x + k)
				flag[x+k] = true
			}
		}
	}
	return ans
}

