func maximumSum(arr []int) int {
	l := len(arr)
	left := make([]int, l)
	right := make([]int, l)

	sum := 0
	for i := 0; i < l-1; i++ {
		sum += arr[i]
		left[i] = sum
		if sum < 0 {
			sum = 0
		}
	}

	sum = 0
	for i := l - 1; i >= 1; i-- {
		sum += arr[i]
		right[i] = sum
		if sum < 0 {
			sum = 0
		}
	}

	ans := arr[0]
	sum = 0
	for _, a := range arr {
		sum += a
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	if len(arr) == 1 {
		return ans
	}

	var tmp int
	for i := 0; i < l; i++ {
		tmp = 0
		if i > 0 {
			tmp += left[i-1]
		}
		if i < l-1 {
			tmp += right[i+1]
		}
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}

