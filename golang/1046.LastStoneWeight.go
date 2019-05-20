type Heap struct {
	ele []int
}

func (h Heap) Len() int {
	return len(h.ele)
}

func (h Heap) Less(i, j int) bool {
	return h.ele[i] > h.ele[j]
}

func (h Heap) Swap(i, j int) {
	h.ele[i], h.ele[j] = h.ele[j], h.ele[i]
}

func (h *Heap) Push(x interface{}) {
	h.ele = append(h.ele, x.(int))
}

func (h *Heap) Pop() interface{} {
	l := len(h.ele) - 1
	e := h.ele[l]
	h.ele = h.ele[:l]
	return e
}

func lastStoneWeight(stones []int) int {
	var h Heap
	for _, e := range stones {
		heap.Push(&h, e)
	}

	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x == y {
			continue
		}
		if x > y {
			x, y = y, x
		}
		heap.Push(&h, y-x)
	}

	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(&h).(int)
}

