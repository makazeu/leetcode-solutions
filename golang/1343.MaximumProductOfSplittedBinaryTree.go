/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const MOD = 1000000007

func maxProduct(root *TreeNode) int {
    var sums []int
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftSum := dfs(node.Left)
        if leftSum != 0 {
            sums = append(sums, leftSum)
        }
        
        rightSum := dfs(node.Right)
        if rightSum != 0 {
            sums = append(sums, rightSum)
        }

        return node.Val + leftSum + rightSum
    }
    
    sum := dfs(root)
    halfSum := sum / 2
    minDiff := math.MaxInt32
    var ans int
    
    for _, x := range sums {
        diff := abs(halfSum - x)
        if diff < minDiff {
            minDiff = diff
            ans = int((int64(x) * int64(sum - x)) % int64(MOD))
        }
        if x >= halfSum {
            break
        }
    }
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

