/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var nums []int
	f := false
	for head != nil {
		if head.Val == 1 {
			f = true
		}
		if f {
			nums = append(nums, head.Val)
		}

		head = head.Next
	}
	if len(nums) == 0 {
		return 0
	}

	prod := 0
	x := 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod += x * nums[i]
		x *= 2
	}
	return prod
}

