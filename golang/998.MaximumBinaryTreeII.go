/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root.Val > val {
		if root.Right == nil {
			root.Right = new(TreeNode)
			root.Right.Val = val
		} else {
			root.Right = insertIntoMaxTree(root.Right, val)
		}
		return root
	} else {
		node := new(TreeNode)
		node.Val = val
		node.Left = root
		return node
	}
}

