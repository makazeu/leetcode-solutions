/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	m map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	m := make(map[int]bool)
	var dfs func(node *TreeNode, val int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		m[val] = true
		dfs(node.Left, 2*val+1)
		dfs(node.Right, 2*val+2)
	}
	dfs(root, 0)
	return FindElements{m: m}
}

func (this *FindElements) Find(target int) bool {
	return this.m[target]
}

