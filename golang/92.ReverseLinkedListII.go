/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    root := new(ListNode)
    root.Next = head
    cnt := root
    
    for i:=1; i<m; i++ {
        cnt = cnt.Next
    }
    
    head = cnt.Next
    cnt.Next = nil
    
    for i:=m; i<=n; i++ {
        tmp := head
        head = head.Next
        tmp.Next = cnt.Next
        cnt.Next = tmp
    }
    
    for cnt.Next != nil {
        cnt = cnt.Next
    }
    cnt.Next = head
    
    return root.Next
}

