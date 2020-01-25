var tmp []int

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	p, q := start, mid+1
	r := start
	for p <= mid && q <= end {
		if nums[p] <= nums[q] {
			tmp[r] = nums[p]
			p++
		} else {
			tmp[r] = nums[q]
			q++
		}
		r++
	}

	for p <= mid {
		tmp[r] = nums[p]
		p++
		r++
	}
	for q <= end {
		tmp[r] = nums[q]
		q++
		r++
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i]
	}
}

func sortArray(nums []int) []int {
	tmp = make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

