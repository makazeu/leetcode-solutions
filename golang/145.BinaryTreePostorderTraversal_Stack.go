const (
	First = iota
	Left
	Right
)

type Item struct {
	node   *TreeNode
	status int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}

	stack := list.New()
	stack.PushFront(&Item{root, First})

	for stack.Len() > 0 {
		item := stack.Front()
		node, status := item.Value.(*Item).node, item.Value.(*Item).status

		if status == First {
			if node.Left != nil {
				item.Value.(*Item).status = Left
				stack.PushFront(&Item{node.Left, First})
				continue
			}
		}
		if status != Right {
			if node.Right != nil {
				item.Value.(*Item).status = Right
				stack.PushFront(&Item{node.Right, First})
				continue
			}
		}
		nums = append(nums, node.Val)
		stack.Remove(item)
	}

	return nums
}

