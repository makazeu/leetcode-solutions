/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := new(ListNode)
	now := list
	var a, b, c int
	for {
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		} else {
			a = 0
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		} else {
			b = 0
		}
		s := a + b + c
		c = s / 10
		a = s % 10

		now.Val = a

		if l1 == nil && l2 == nil {
			break
		} else {
			now.Next = new(ListNode)
			now = now.Next
		}
	}
	if c > 0 {
		now.Next = new(ListNode)
		now = now.Next
		now.Val = c
	}
	return list
}
