func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	exponentNum := strings.Count(s, "e")
	if exponentNum > 1 {
		return false
	}
	if exponentNum == 0 {
		return validate(s)
	}

	exponentPos := strings.Index(s, "e")
	if exponentPos == 0 || exponentPos == len(s)-1 {
		return false
	}

	beforeE := s[:exponentPos]
	if !validate(beforeE) {
		return false
	}
	afterE := s[exponentPos+1:]
	if strings.Count(afterE, ".") > 0 {
		return false
	}
	return validate(afterE)
}

func validate(s string) bool {
	_, err := strconv.ParseFloat(s, 10)
	if err == nil {
		return true
	}
	return false
}
