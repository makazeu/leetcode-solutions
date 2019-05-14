/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var ans []int
	if root != nil {
		dfs(root, 1, &ans)
	}

	return ans
}

func dfs(node *TreeNode, depth int, ans *[]int) {
	if depth > len(*ans) {
		*ans = append(*ans, node.Val)
	} else {
		if (*ans)[depth-1] < node.Val {
			(*ans)[depth-1] = node.Val
		}
	}

	if node.Left != nil {
		dfs(node.Left, depth+1, ans)
	}
	if node.Right != nil {
		dfs(node.Right, depth+1, ans)
	}
}

