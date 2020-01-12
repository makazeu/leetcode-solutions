/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var res []int
	traversal(root1, &res)
	traversal(root2, &res)
	sort.Ints(res)
	return res
}

func traversal(node *TreeNode, array *[]int) {
	if node == nil {
		return
	}
	*array = append(*array, node.Val)
	traversal(node.Left, array)
	traversal(node.Right, array)
}

