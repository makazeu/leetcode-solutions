type Items struct {
	h    map[int]int
	nums []int
}

func (items Items) Len() int {
	return len(items.nums)
}

func (items Items) Less(i, j int) bool {
	a := items.h[items.nums[i]]
	b := items.h[items.nums[j]]

	if a == 0 && b == 0 {
		return items.nums[i] < items.nums[j]
	}
	if a == 0 && b != 0 {
		return false
	}
	if b == 0 && a != 0 {
		return true
	}
	return a < b
}

func (items Items) Swap(i, j int) {
	items.nums[i], items.nums[j] = items.nums[j], items.nums[i]
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var items Items
	items.h = make(map[int]int)
	for i, e := range arr2 {
		items.h[e] = i + 1
	}
	items.nums = arr1
	sort.Sort(items)

	return items.nums
}

