/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MaxInt32
	dfs(root, 0, &ans)
	return ans
}

func dfs(node *TreeNode, depth int, ans *int) {
	depth++
	if node.Right == nil && node.Left == nil {
		*ans = min(*ans, depth)
		return
	}
	if node.Left != nil {
		dfs(node.Left, depth, ans)
	}
	if node.Right != nil {
		dfs(node.Right, depth, ans)
	}
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

