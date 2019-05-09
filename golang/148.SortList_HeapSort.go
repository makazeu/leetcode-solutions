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
func sortList(head *ListNode) *ListNode {
	header := new(ListNode)

	h := new(Heap)
	for ; head != nil; head = head.Next {
		h.Push(head)
	}
	heap.Init(h)

	cur := header
	for {
		if h.Len() == 0 {
			break
		}

		cur.Next = heap.Pop(h).(*ListNode)
		cur = cur.Next
	}
	cur.Next = nil

	return header.Next
}

