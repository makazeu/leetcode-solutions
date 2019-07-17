/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var ans *TreeNode
	var lowest int
	dfs(root, 0, &lowest, &ans)
	return ans
}

func dfs(node *TreeNode, depth int, lowest *int, ans **TreeNode) int {
	lowestDepth := depth
	leftDepth := -1
	if node.Left != nil {
		leftDepth = dfs(node.Left, depth+1, lowest, ans)
		if leftDepth > lowestDepth {
			lowestDepth = leftDepth
		}
	}
	rightDepth := -1
	if node.Right != nil {
		rightDepth = dfs(node.Right, depth+1, lowest, ans)
		if rightDepth > lowestDepth {
			lowestDepth = rightDepth
		}
	}

	if lowestDepth > *lowest {
		*lowest = lowestDepth
	}

	if leftDepth == rightDepth && *lowest == lowestDepth {
		*ans = node
	}
	return lowestDepth
}

