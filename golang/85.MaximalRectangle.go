type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Push(e interface{}) {
	stack.list.PushFront(e)
}

func (stack *Stack) Top() interface{} {
	if stack.Size() == 0 {
		return nil
	}
	return stack.list.Front().Value
}

func (stack *Stack) Pop() interface{} {
	if stack.Size() == 0 {
		return nil
	}
	e := stack.list.Front()
	stack.list.Remove(e)
	return e.Value
}

func (stack *Stack) Size() int {
	return stack.list.Len()
}

type Node struct {
	height, num int
}

func calc(heights []int) int {
	stack := NewStack()
	heights = append(heights, 0)
	max := 0

	for _, h := range heights {
		if stack.Size() == 0 || stack.Top().(*Node).height <= h {
			stack.Push(&Node{h, 1})
			continue
		}

		num := 0
		for {
			if stack.Size() == 0 || stack.Top().(*Node).height <= h {
				break
			}
			top := stack.Pop().(*Node)
			num += top.num

			if num*top.height > max {
				max = num * top.height
			}
		}
		stack.Push(&Node{h, num + 1})
	}

	return max
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	max := 0

	heights := make([][]int, 2)
	for i := range heights {
		heights[i] = make([]int, n)
	}
	for i := range matrix {
		for j, x := range matrix[i] {
			if x == '0' {
				heights[i%2][j] = 0
				continue
			}
			heights[i%2][j] = 1
			if i > 0 {
				heights[i%2][j] += heights[(i-1)%2][j]
			}
		}

		cur := calc(heights[i%2])
		if cur > max {
			max = cur
		}
	}

	return max
}

