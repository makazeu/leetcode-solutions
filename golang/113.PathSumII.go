/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(cur *TreeNode, depth, left int, ans []int, answers *[][]int) {
	if cur == nil {
		return
	}
	left -= cur.Val
	ans = append(ans, cur.Val)
	if cur.Left == nil && cur.Right == nil {
		if left == 0 {
			tmp := make([]int, depth+1)
			copy(tmp, ans)
			*answers = append(*answers, tmp)
		}
	}

	if cur.Left != nil {
		dfs(cur.Left, depth+1, left, ans, answers)
	}
	if cur.Right != nil {
		dfs(cur.Right, depth+1, left, ans, answers)
	}
	left += cur.Val
	ans = ans[:depth]
}

func pathSum(root *TreeNode, sum int) [][]int {
	var ans []int
	var answers [][]int
	dfs(root, 0, sum, ans, &answers)

	return answers
}
