func sortArray(nums []int) []int {
    l := len(nums)
	for i := l/2 - 1; i >= 0; i-- {
		heap(nums, i, l-1)
	}

	for i := l - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heap(nums, 0, i-1)
	}
    
    return nums
}

func heap(arr []int, i, end int) {
	parent := i
	child := parent*2 + 1
	for child <= end {
		if child+1 <= end && arr[child] < arr[child+1] {
			child++
		}
		if arr[parent] >= arr[child] {
			return
		}
		arr[parent], arr[child] = arr[child], arr[parent]
		parent = child
		child = parent*2 + 1
	}
}

