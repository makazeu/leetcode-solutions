/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	root := new(ListNode)
	root.Next = head
	fast, slow := root, root

	l := 0
	for ; fast != nil; l++ {
		fast = fast.Next
		if l%2 == 1 {
			slow = slow.Next
		}
	}
	l--

	if l <= 1 {
		return true
	}

	mid := new(ListNode)
	slow = slow.Next
	for slow != nil {
		tmp := slow
		slow = slow.Next
		tmp.Next = mid.Next
		mid.Next = tmp
	}

	for mid != nil {
		if mid.Val != root.Val {
			return false
		}
		mid = mid.Next
		root = root.Next
	}
	return true
}

