type Event struct {
	t, e int
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	a := make([]*Event, len(A)*2)
	for i, e := range A {
		a[i*2] = &Event{e[0], 1}
		a[i*2+1] = &Event{e[1], -1}
	}
	b := make([]*Event, len(B)*2)
	for i, e := range B {
		b[i*2] = &Event{e[0], 1}
		b[i*2+1] = &Event{e[1], -1}
	}

	var ans [][]int
	var p, q int
	var x, y *Event
	var status, t int
	start := -1
	for p < len(a) && q < len(b) {
		x, y = a[p], b[q]
		if compareEvent(x, y) {
			status += x.e
			t = x.t
			p++
		} else {
			status += y.e
			t = y.t
			q++
		}

		if status == 2 && start == -1 {
			start = t
		} else if status == 1 && start != -1 {
			ans = append(ans, []int{start, t})
			start = -1
		}
	}
	return ans
}

func compareEvent(e1, e2 *Event) bool {
	if e1.t == e2.t {
		return e1.e > e2.e
	}
	return e1.t < e2.t
}

