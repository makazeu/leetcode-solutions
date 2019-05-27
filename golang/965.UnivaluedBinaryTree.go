/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root, root.Val)
}

func dfs(t *TreeNode, val int) bool {
	if t == nil {
		return true
	}
	if val != t.Val {
		return false
	}
	return dfs(t.Right, val) && dfs(t.Left, val)
}

