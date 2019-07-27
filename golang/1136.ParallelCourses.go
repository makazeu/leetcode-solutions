func minimumSemesters(N int, relations [][]int) int {
	pre := make([]int, N+1)
	nxt := make([][]int, N+1)
	for _, e := range relations {
		nxt[e[0]] = append(nxt[e[0]], e[1])
		pre[e[1]]++
	}

	lst := list.New()
	for i := 1; i <= N; i++ {
		lst.PushBack(i)
	}

	ans := 0
	for lst.Len() > 0 {
		f := false
		n := lst.Front()
		h := make(map[int]int)
		for n != nil {
			e := n
			n = n.Next()
			c := e.Value.(int)
			if pre[c] > 0 {
				continue
			}
			f = true
			lst.Remove(e)

			for _, nc := range nxt[c] {
				h[nc]++
			}
		}
		ans++
		if !f {
			return -1
		}
		for k, v := range h {
			pre[k] -= v
		}
	}
	return ans
}

