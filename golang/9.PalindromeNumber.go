func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	var slice []int
	for ; x > 0; x /= 10 {
		slice = append(slice, x%10)
	}

	length := len(slice)
	for i := 0; i < length; i++ {
		if slice[i] != slice[length-i-1] {
			return false
		}
	}
	return true
}
