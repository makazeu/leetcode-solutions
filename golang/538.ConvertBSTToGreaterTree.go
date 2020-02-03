/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    solve(root, 0)
    return root
}

func solve(node *TreeNode, greater int) int {
    if node == nil {
        return greater
    }
    greater = solve(node.Right, greater)
    greater += node.Val
    node.Val = greater
    greater = solve(node.Left, greater)
    return greater
}

