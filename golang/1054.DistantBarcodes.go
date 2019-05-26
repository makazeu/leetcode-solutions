type Item struct {
	num, count int
}

type Heap struct {
	items []*Item
}

func (h Heap) Len() int {
	return len(h.items)
}

func (h Heap) Less(i, j int) bool {
	return h.items[i].count > h.items[j].count
}

func (h Heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap) Pop() interface{} {
	l := h.Len() - 1
	e := h.items[l]
	h.items = h.items[:l]
	return e
}

func (h *Heap) Push(x interface{}) {
	h.items = append(h.items, x.(*Item))
}

func rearrangeBarcodes(barcodes []int) []int {
	var h Heap
	m := make(map[int]int)
	for _, e := range barcodes {
		m[e]++
	}
	for k, v := range m {
		heap.Push(&h, &Item{k, v})
	}

	last := 0
	var ans []int
	for h.Len() > 0 {
		e1 := heap.Pop(&h).(*Item)
		if e1.num != last {
			ans = append(ans, e1.num)
			last = e1.num
			e1.count--
			if e1.count != 0 {
				heap.Push(&h, e1)
			}
			continue
		}

		e2 := heap.Pop(&h).(*Item)
		ans = append(ans, e2.num)
		last = e2.num
		e2.count--
		if e2.count != 0 {
			heap.Push(&h, e2)
		}
		heap.Push(&h, e1)
	}
	return ans
}

