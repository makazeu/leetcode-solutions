/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = new(TreeNode)
			root.Left.Val = val
		} else {
			insertIntoBST(root.Left, val)
		}
	} else if val > root.Val {
		if root.Right == nil {
			root.Right = new(TreeNode)
			root.Right.Val = val
		} else {
			insertIntoBST(root.Right, val)
		}
	}
	return root
}

