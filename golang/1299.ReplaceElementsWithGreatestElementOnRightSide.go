func replaceElements(arr []int) []int {
	m := -1
	for i := len(arr) - 1; i >= 0; i-- {
		x := arr[i]
		arr[i] = m
		if x > m {
			m = x
		}
	}
	return arr
}

