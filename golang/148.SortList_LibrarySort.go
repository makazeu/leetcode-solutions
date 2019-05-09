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
		*h = append(*h, head)
	}
	sort.Sort(h)

	cur := header
	for _, node := range *h {
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = nil

	return header.Next
}

