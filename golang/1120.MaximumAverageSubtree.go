/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maximumAverageSubtree(root *TreeNode) float64 {
	ans := float64(0)
	dfs(root, &ans)
	return ans
}

func dfs(node *TreeNode, ans *float64) (int, int) {
	if node == nil {
		return 0, 0
	}
	sum := node.Val
	num := 1

	a, b := dfs(node.Left, ans)
	sum += a
	num += b
	a, b = dfs(node.Right, ans)
	sum += a
	num += b

	v := float64(sum) / float64(num)
	if v > *ans {
		*ans = v
	}
	return sum, num
}

