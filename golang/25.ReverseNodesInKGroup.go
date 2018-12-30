/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func peekListLength(list *ListNode, k int) bool {
	num := 0
	for list != nil {
		num++
		if num >= k {
			return true
		}
		list = list.Next
	}
	return false
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	latest := new(ListNode)
	first := latest
	for true {
		if head == nil {
			break
		}
		if peekListLength(head, k) {
			tail := head
			head = head.Next
			tail.Next = nil
			cur := tail
			for i := 0; i < k-1; i++ {
				next := head.Next
				head.Next = cur
				cur = head
				head = next
			}
			latest.Next = cur
			latest = tail
		} else {
			for head != nil {
				latest.Next = head
				latest = latest.Next
				head = head.Next
			}
		}
	}
	return first.Next
}
