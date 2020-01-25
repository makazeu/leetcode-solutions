/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	myRoot := new(TreeNode)
	myRoot.Right = root
	deleteLeafNode(root, myRoot, target)
	return myRoot.Right
}

func deleteLeafNode(node, parent *TreeNode, target int) {
	if node == nil {
		return
	}
	deleteLeafNode(node.Left, node, target)
	deleteLeafNode(node.Right, node, target)

	if node.Val != target {
		return
	}
	if node.Left != nil || node.Right != nil {
		return
	}

	if parent.Left == node {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
}

