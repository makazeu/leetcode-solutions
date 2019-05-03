/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	return dfs(root, &k)
}

func dfs(node *TreeNode, k *int) (ans int) {
	if node.Left != nil {
		if ans = dfs(node.Left, k); *k == 0 {
			return
		}
	}

	if *k--; *k == 0 {
		return node.Val
	}

	if node.Right != nil {
		ans = dfs(node.Right, k)
	}
	return
}

