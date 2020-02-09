tion for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    
    root := new(ListNode)
    root.Next = head
    
    cnt := head.Next
    pre := head
    for cnt != nil {
        if cnt.Val >= pre.Val {
            pre = cnt
            cnt = cnt.Next
            continue
        }
        
        pre.Next = cnt.Next
        pos := root
        for pos.Next.Val <= cnt.Val {
            pos = pos.Next
        }
        cnt.Next = pos.Next
        pos.Next = cnt
        cnt = pre.Next
    }
    return root.Next
}

