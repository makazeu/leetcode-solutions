/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    root := new(ListNode)
    
    for head != nil {
        tmp := head
        head = head.Next
        tmp.Next = root.Next
        root.Next = tmp
    }
    
    return root.Next
}

