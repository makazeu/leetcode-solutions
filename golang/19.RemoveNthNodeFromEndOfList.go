/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	header := new(ListNode)
	header.Next = head

	front, back := head, head
	backPrev := header
	for i := 0; i < n; i++ {
		front = front.Next
	}

	for {
		if front == nil {
			break
		}
		front = front.Next
		back = back.Next
		backPrev = backPrev.Next
	}
	backPrev.Next = backPrev.Next.Next
	return header.Next
}

