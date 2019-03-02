/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	dx, px := findNode(root, x, 0, nil)
	dy, py := findNode(root, y, 0, nil)

	return dx == dy && px != py
}

func findNode(root *TreeNode, x, depth int, parent *TreeNode) (int, *TreeNode) {
	if root.Val == x {
		return depth, parent
	}
	if root.Left != nil {
		retDepth, retTreeNode := findNode(root.Left, x, depth+1, root)
		if retDepth != -1 {
			return retDepth, retTreeNode
		}
	}
	if root.Right != nil {
		retDepth, retTreeNode := findNode(root.Right, x, depth+1, root)
		if retDepth != -1 {
			return retDepth, retTreeNode
		}
	}
	return -1, nil
}

