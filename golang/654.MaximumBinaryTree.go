/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := new(TreeNode)
	buildTree(root, nums)
	return root
}

func buildTree(tree *TreeNode, nums []int) {
	if len(nums) == 1 {
		tree.Val = nums[0]
		return
	}

	maxPos := findMaxElementPosition(nums)
	tree.Val = nums[maxPos]

	if maxPos > 0 {
		tree.Left = new(TreeNode)
		buildTree(tree.Left, nums[0:maxPos])
	}
	if maxPos+1 < len(nums) {
		tree.Right = new(TreeNode)
		buildTree(tree.Right, nums[maxPos+1:])
	}
}

func findMaxElementPosition(nums []int) int {
	pos := 0
	for i, x := range nums {
		if x > nums[pos] {
			pos = i
		}
	}
	return pos
}
