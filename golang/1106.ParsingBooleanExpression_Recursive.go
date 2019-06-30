func parseBoolExpr(expression string) bool {
	if strings.EqualFold(expression, "t") {
		return true
	}
	if strings.EqualFold(expression, "f") {
		return false
	}
	if strings.HasPrefix(expression, "!") {
		return !parseBoolExpr(expression[2:len(expression)-1])
	}
	if strings.HasPrefix(expression, "&") {
		deep := 0
		start := 2
		for i := 2; i < len(expression)-1; i++ {
			if expression[i] == ',' && deep == 0 {
				if !parseBoolExpr(expression[start:i]) {
					return false
				}
				start = i + 1
				continue
			}
			if expression[i] == '(' {
				deep++
			} else if expression[i] == ')' {
				deep--
			}
		}
		return parseBoolExpr(expression[start : len(expression)-1])
	} else if strings.HasPrefix(expression, "|") {
		deep := 0
		start := 2
		for i := 2; i < len(expression)-1; i++ {
			if expression[i] == ',' && deep == 0 {
				if parseBoolExpr(expression[start:i]) {
					return true
				}
				start = i + 1
				continue
			}
			if expression[i] == '(' {
				deep++
			} else if expression[i] == ')' {
				deep--
			}
		}
		return parseBoolExpr(expression[start : len(expression)-1])
	} else {
		return true
	}
}

