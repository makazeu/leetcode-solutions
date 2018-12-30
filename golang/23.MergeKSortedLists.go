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

	for true {
		minVal := int(^uint(0) >> 1)
		minPos := -1

		for pos, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < minVal {
				minVal = list.Val
				minPos = pos
			}
		}
		if minPos == -1 {
			break
		}

		cur := lists[minPos]
		lists[minPos] = cur.Next
		cur.Next = nil
		prev.Next = cur
		prev = cur
	}

	return head.Next
}
