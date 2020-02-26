/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var ans []int
    ans = append(ans, root.Val)
    if root.Left == nil && root.Right == nil {
        return ans
    }

    var lefts []int
    if root.Left != nil {
        findLeftBoundary(root.Left, &lefts)
    }
    fmt.Println(lefts)

    var rights []int
    if root.Right != nil {
        findRightBoundary(root.Right, &rights)
    }
    fmt.Println(rights)

    var leaves []int
    findLeaves(root, &leaves)

    if len(lefts) > 0 {
        leaves = leaves[1:]
    }
    if len(rights) > 0 {
        leaves = leaves[:len(leaves) - 1]
    }
    fmt.Println(leaves)

    ans = append(ans, lefts...)
    ans = append(ans, leaves...)
    ans = append(ans, rights...)
    return ans
}

func findLeaves(node *TreeNode, nums *[]int) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *nums = append(*nums, node.Val)
        return
    }
    findLeaves(node.Left, nums)
    findLeaves(node.Right, nums)
}

func findLeftBoundary(node *TreeNode, nums *[]int) {
    *nums = append(*nums, node.Val)
    if node.Left != nil {
        findLeftBoundary(node.Left, nums)
    } else if node.Right != nil {
        findLeftBoundary(node.Right, nums)
    }
}

func findRightBoundary(node *TreeNode, nums *[]int) {
    if node.Right != nil {
        findRightBoundary(node.Right, nums)
    } else if node.Left != nil {
        findRightBoundary(node.Left, nums)
    }
    *nums = append(*nums, node.Val)
}

