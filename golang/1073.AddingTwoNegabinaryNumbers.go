func addNegabinary(arr1 []int, arr2 []int) []int {
	var res []int
	var c int
	for i := 0; i < len(arr1) + 2 || i < len(arr2) + 2; i++ {
		var b1, b2 int
		if len(arr1)-i-1 >= 0 {
			b1 = arr1[len(arr1)-i-1]
		}
		if len(arr2)-i-1 >= 0 {
			b2 = arr2[len(arr2)-i-1]
		}
		r := c
		if i%2 == 0 {
			r += b1 + b2
		} else {
			r -= b1 + b2
		}

		if i%2 == 0 {
			if r == 3 {
				r = 1
				c = 1
			} else if r == 2 {
				r = 0
				c = 1
			} else if r == -1 {
				r = 1
				c = -1
			} else {
				c = 0
			}
		} else {
			if r == -3 {
				r = 1
				c = -1
			} else if r == -2 {
				r = 0
				c = -1
			} else if r == -1 {
				r = 1
				c = 0
			} else if r == 1 {
				r = 1
				c = 1
			} else {
				c = 0
			}
		}

		res = append(res, r)
	}

	var ans []int
	pos := len(res) - 1
	for pos >= 0 && res[pos] == 0 {
		pos--
	}

	for ; pos >= 0; pos-- {
		ans = append(ans, res[pos])
	}

	if len(ans) == 0 {
		ans = append(ans, 0)
	}
	return ans
}

