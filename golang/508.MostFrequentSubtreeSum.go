/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	max := 0
	var ans []int
	m := make(map[int]int)

	if root == nil {
		return ans
	}

	dfs(root, &m, &max, &ans)
	return ans
}

func dfs(node *TreeNode, m *map[int]int, max *int, ans *[]int) int {
	sum := node.Val

	if node.Left != nil {
		sum += dfs(node.Left, m, max, ans)
	}
	if node.Right != nil {
		sum += dfs(node.Right, m, max, ans)
	}

	n := (*m)[sum] + 1
	(*m)[sum] = n
	if n == *max {
		*ans = append(*ans, sum)
	} else if n > *max {
		*max = n
		*ans = []int{sum}
	}

	return sum
}

