/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	my(root, preorder)
	return root
}

func my(node *TreeNode, nums []int) {
	node.Val = nums[0]

	pos := 1
	for ; pos < len(nums); pos++ {
		if nums[pos] > node.Val {
			break
		}
	}
	if pos > 1 {
		node.Left = new(TreeNode)
		my(node.Left, nums[1:pos])
	}
	if pos < len(nums) {
		node.Right = new(TreeNode)
		my(node.Right, nums[pos:])
	}
}

