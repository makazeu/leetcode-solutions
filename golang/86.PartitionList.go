/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	heada := new(ListNode)
	headb := new(ListNode)

	nowa, nowb := heada, headb

	for ; head != nil; {
		if head.Val < x {
			nowa.Next = head
			head = head.Next
			nowa = nowa.Next
			nowa.Next = nil
		} else {
			nowb.Next = head
			head = head.Next
			nowb = nowb.Next
			nowb.Next = nil
		}
	}
	nowa.Next = headb.Next
	return heada.Next
}

