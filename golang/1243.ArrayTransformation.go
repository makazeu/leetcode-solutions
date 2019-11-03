func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return arr
	}

	for {
		var incr []int
		var decr []int
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				incr = append(incr, i)
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				decr = append(decr, i)
			}
		}

		if len(incr) == 0 && len(decr) == 0 {
			break
		}
		for _, i := range incr {
			arr[i]++
		}
		for _, i := range decr {
			arr[i]--
		}
	}
	return arr
}

