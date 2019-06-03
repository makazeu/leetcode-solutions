func check(str, div string) bool {
	for now := div; !strings.EqualFold(str, now); now = now + div {
		if len(now) > len(str) {
			return false
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	l := len(str2)
	for i := 1; i <= l; i++ {
		if l%i != 0 {
			continue
		}
		div := str2[:l/i]
		if check(str2, div) && check(str1, div) {
			return div
		}
	}
	return ""
}

