type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Push(v interface{}) {
	stack.list.PushBack(v)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func longestValidParentheses(s string) int {
	stack := NewStack()

	max, now := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			stack.Push(now)
			now = 0
			break
		case ')':
			if e := stack.Pop(); e != nil {
				now = e.(int) + now + 2
				max = Max(max, now)
			} else {
				max = Max(max, now)
				now = 0
			}
		}
	}
	return Max(max, now)
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
