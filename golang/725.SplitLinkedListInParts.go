/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	total := 0
	head := new(ListNode)
	head.Next = root
	for head.Next != nil {
		head = head.Next
		total++
	}
	num := total / k
	nums := make([]int, k)
	for i := 0; i < k; i++ {
		nums[i] = num
	}
	for i := 0; i < total-num*k; i++ {
		nums[i]++
	}

	var ans []*ListNode
	for _, n := range nums {
		if n == 0 {
			ans = append(ans, nil)
			continue
		}
		h := root
		for i := 1; i <= n-1; i++ {
			root = root.Next
		}
		tmp := root.Next
		root.Next = nil
		root = tmp
		ans = append(ans, h)
	}
	return ans
}

