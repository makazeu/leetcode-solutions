/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	sum := 0
	dfs(root, &sum, -1, -1)
	return sum
}

func dfs(node *TreeNode, sum *int, parent, grandparent int) {
	if node == nil {
		return
	}
	if grandparent%2 == 0 {
		*sum += node.Val
	}
	dfs(node.Left, sum, node.Val, parent)
	dfs(node.Right, sum, node.Val, parent)
}

