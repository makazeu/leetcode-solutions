func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var ans [][]int
	mind := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < mind {
			mind = arr[i+1] - arr[i]
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == mind {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}

