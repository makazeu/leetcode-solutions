func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lg, sm := -1, -1
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			lg = mid
		}
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	l, r = 0, len(nums)-1
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			sm = mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return []int{sm, lg}
}

