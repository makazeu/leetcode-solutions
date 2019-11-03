/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	header := new(ListNode)
	header.Next = head
	flag := true
	for flag {
		if header.Next == nil {
			break
		}

		flag = false
		sum := 0
		m := make(map[int]*ListNode)
		m[0] = header
		node := header.Next
		for node != nil {
			sum += node.Val
			if pre, exists := m[sum]; exists {
				pre.Next = node.Next
				flag = true
				break
			} else {
				m[sum] = node
				node = node.Next
			}
		}
	}

	return header.Next
}

