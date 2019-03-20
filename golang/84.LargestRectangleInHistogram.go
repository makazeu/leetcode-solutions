type Node struct {
	height, num int
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	max := 0

	stack := list.New()
	for _, h := range heights {
		if stack.Len() == 0 || stack.Front().Value.(Node).height <= h {
			stack.PushFront(Node{h, 1})
			continue
		}

		num := 0
		for {
			if stack.Len() == 0 {
				break
			}

			node := stack.Front().Value.(Node)
			if node.height <= h {
				break
			}
			stack.Remove(stack.Front())

			num += node.num
			if num*node.height > max {
				max = num * node.height
			}
		}
		stack.PushFront(Node{h, num + 1})
	}

	return max
}

