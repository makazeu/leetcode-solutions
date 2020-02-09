/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        
        if left < 0 {
            left = 0
        }
        if right < 0 {
            right = 0
        }
        result := left + node.Val + right
        if result > ans {
            ans = result
        }
        
        if left >= right {
            return node.Val + left
        }
        return node.Val + right
    }
    dfs(root)
    return ans
}

