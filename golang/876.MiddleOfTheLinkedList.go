/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	header := new(ListNode)
	header.Next = head
	a, b := header, header
	for {
		if b == nil {
			return a
		}
		b = b.Next
		a = a.Next
		if b != nil {
			b = b.Next
		}
	}
}

