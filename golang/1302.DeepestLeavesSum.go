/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	maxDepth := getMaxDepth(root, 0)
	return sumTree(root, 0, maxDepth)
}

func sumTree(node *TreeNode, prevDepth, maxDepth int) int {
	if node == nil {
		return 0
	}
	if prevDepth+1 == maxDepth {
		return node.Val
	}
	return sumTree(node.Left, prevDepth+1, maxDepth) + sumTree(node.Right, prevDepth+1, maxDepth)
}

func getMaxDepth(node *TreeNode, prevDepth int) int {
	currentDepth := prevDepth + 1
	if node.Left != nil {
		currentDepth = max(currentDepth, getMaxDepth(node.Left, prevDepth+1))
	}
	if node.Right != nil {
		currentDepth = max(currentDepth, getMaxDepth(node.Right, prevDepth+1))
	}
	return currentDepth
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

