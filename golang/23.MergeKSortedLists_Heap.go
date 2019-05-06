type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	l := h.Len()
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head := new(ListNode)
	prev := head

	h := new(Heap)
	for _, list := range lists {
		if list != nil {
			h.Push(list)
		}
	}
	heap.Init(h)

	for {
		if h.Len() == 0 {
			break
		}
		
		cur := (*h)[0]
		if cur.Next == nil {
			heap.Pop(h)
		} else {
			(*h)[0] = cur.Next
			heap.Fix(h, 0)
		}

		cur.Next = nil
		prev.Next = cur
		prev = cur
	}

	return head.Next
}

