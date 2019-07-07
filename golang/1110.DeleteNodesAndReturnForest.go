/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	hash := make(map[int]bool)
	for _, e := range to_delete {
		hash[e] = true
	}
	var ans []*TreeNode
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		node := q.Front().Value.(*TreeNode)
		q.Remove(q.Front())
		if !deleteNode(node, nil, q, hash) {
			ans = append(ans, node)
		}
	}
	return ans
}

func deleteNode(node, parent *TreeNode, q *list.List, hash map[int]bool) bool {
	if node == nil {
		return true
	}
	if hash[node.Val] {
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		if node.Left != nil {
			q.PushBack(node.Left)
		}
		if node.Right != nil {
			q.PushBack(node.Right)
		}
		return true
	} else {
		deleteNode(node.Left, node, q, hash)
		deleteNode(node.Right, node, q, hash)
		return false
	}
}

