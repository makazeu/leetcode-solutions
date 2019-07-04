/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	sums := make(map[int]int)
	nums := make(map[int]int)
	travel(root, sums, nums, 0)
	level := 0
	var ans []float64
	for {
		sum, exists := sums[level]
		if !exists {
			break
		}
		ans = append(ans, float64(sum)/float64(nums[level]))
		level++
	}
	return ans
}

func travel(node *TreeNode, sums, nums map[int]int, level int) {
	if node == nil {
		return
	}
	sums[level] += node.Val
	nums[level] ++
	travel(node.Left, sums, nums, level+1)
	travel(node.Right, sums, nums, level+1)
}

