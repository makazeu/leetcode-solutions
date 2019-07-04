type Item struct {
	node  *TreeNode
	level int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	queue := list.New()
	queue.PushBack(Item{root, 0})
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		item := front.Value.(Item)
		node, level := item.node, item.level
		for len(ans) <= level {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], node.Val)

		if node.Left != nil {
			queue.PushBack(Item{node.Left, level + 1})
		}
		if node.Right != nil {
			queue.PushBack(Item{node.Right, level + 1})
		}
	}

	for i := range ans {
		if i%2 == 0 {
			continue
		}
		l := len(ans[i])
		for j := 0; j < l/2; j++ {
			ans[i][j], ans[i][l-j-1] = ans[i][l-j-1], ans[i][j]
		}
	}
	return ans
}

