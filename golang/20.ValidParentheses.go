func pop(list *list.List) rune {
	if list.Len() == 0 {
		return 0
	}
	e := list.Front()
	list.Remove(e)
	return e.Value.(rune)
}

func isValid(s string) bool {
	stack := list.New()

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.PushFront(c)
		case ')', ']', '}':
			r := pop(stack)
			if c == ')' && r != '(' || c == ']' && r != '[' || c == '}' && r != '{' {
				return false
			}
		}
	}

	return stack.Len() == 0
}
