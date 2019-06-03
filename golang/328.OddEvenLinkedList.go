/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := new(ListNode)
	even := new(ListNode)
	now := head

	oddNow, evenNow := odd, even
	index := 1
	for ; now != nil; index++ {
		if index&1 == 1 {
			oddNow.Next = now
			oddNow = oddNow.Next
			now = oddNow.Next
			oddNow.Next = nil
		} else {
			evenNow.Next = now
			evenNow = evenNow.Next
			now = evenNow.Next
			evenNow.Next = nil
		}
	}
	if odd.Next == nil {
		return even.Next
	} else {
		oddNow.Next = even.Next
		return odd.Next
	}
}

