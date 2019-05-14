type Node struct {
	id, pre int
}

type Heap struct {
	nodes []*Node
	pos   []int
}

func (h Heap) Len() int {
	return len(h.nodes)
}

func (h Heap) Less(i, j int) bool {
	return h.nodes[i].pre < h.nodes[j].pre
}

func (h Heap) Swap(i, j int) {
	h.pos[h.nodes[i].id], h.pos[h.nodes[j].id] = j, i
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *Heap) Pop() interface{} {
	l := len(h.nodes) - 1
	e := h.nodes[l]
	h.nodes = h.nodes[:l]
	h.pos[e.id] = -1
	return e
}

func (h *Heap) Push(x interface{}) {
	node := x.(*Node)
	h.pos = append(h.pos, node.id)
	h.nodes = append(h.nodes, node)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var h Heap
	for i := 0; i < numCourses; i++ {
		node := new(Node)
		node.id = i
		h.Push(node)
	}

	next := make([][]int, numCourses)
	for _, pre := range prerequisites {
		h.nodes[pre[0]].pre++
		next[pre[1]] = append(next[pre[1]], pre[0])
	}
	heap.Init(&h)

	for i := 0; i < numCourses; i++ {
		node := heap.Pop(&h).(*Node)
		if node.pre > 0 {
			return false
		}

		for _, nxt := range next[node.id] {
			h.nodes[h.pos[nxt]].pre--
			heap.Fix(&h, h.pos[nxt])
		}
	}

	return true
}

