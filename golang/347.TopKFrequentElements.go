type Item struct {
	ele int
	num int
}

type SortSlice []*Item

func (slice SortSlice) Len() int {
	return len(slice)
}

func (slice SortSlice) Less(i, j int) bool {
	return slice[i].num > slice[j].num
}

func (slice SortSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var items SortSlice
	for k, v := range m {
		items = append(items, &Item{k, v})
	}
	sort.Sort(items)

	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, items[i].ele)
	}
	return ans
}

