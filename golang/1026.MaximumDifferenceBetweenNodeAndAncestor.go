/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchInSubTree(root *TreeNode, result *int) (int, int) {
	min := root.Val
	max := root.Val
	if root.Left != nil {
		curMin, curMax := searchInSubTree(root.Left, result)
		if curMin < min {
			min = curMin
		}
		if curMax > max {
			max = curMax
		}
	}
	if root.Right != nil {
		curMin, curMax := searchInSubTree(root.Right, result)
		if curMin < min {
			min = curMin
		}
		if curMax > max {
			max = curMax
		}
	}

	resMin := abs(min - root.Val)
	resMax := abs(max - root.Val)
	*result = If(resMin > *result, resMin, *result).(int)
	*result = If(resMax > *result, resMax, *result).(int)

	return min, max
}

func abs(a int) int {
	return If(a < 0, -a, a).(int)
}

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	searchInSubTree(root, &ans)
	return ans
}

