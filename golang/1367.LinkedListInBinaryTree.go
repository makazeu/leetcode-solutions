/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	m := make(map[*TreeNode][]bool)
	var dfs func(node, parent *TreeNode) bool
	dfs = func(node, parent *TreeNode) bool {
		if node == nil {
			return false
		}
		m[node] = make([]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			if node.Val != nums[i] {
				m[node][i] = false
				continue
			}
			if i != 0 && parent != nil {
				m[node][i] = m[parent][i-1]
			} else if i == 0 {
				m[node][i] = true
			}
		}

		if m[node][len(nums)-1] {
			return true
		}
		return dfs(node.Left, node) || dfs(node.Right, node)
	}
	return dfs(root, nil)
}

