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
func levelOrder(root *TreeNode) [][]int {
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
	return ans
}

