/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var num []int
	dfs(root, &num)
	return num
}

func dfs(node *TreeNode, num *[]int){
	if node == nil {
		return
	}
	if node.Left != nil {
		dfs(node.Left, num)
	}

	*num = append(*num, node.Val)

	if node.Right != nil {
		dfs(node.Right, num)
	}
	return
}

