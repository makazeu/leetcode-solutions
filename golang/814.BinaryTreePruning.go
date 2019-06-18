/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
	if root != nil {
		run(root)
	}
	return root
}

func run(node *TreeNode) bool {
	flag := true
	if node.Left != nil && run(node.Left) {
		flag = false
	} else {
		node.Left = nil
	}

	if node.Right != nil && run(node.Right) {
		flag = false
	} else {
		node.Right = nil
	}

	return node.Val != 0 || !flag
}

