type Item struct {
	word string
	num  int
}

type SortSlice []*Item

func (slice SortSlice) Len() int {
	return len(slice)
}

func (slice SortSlice) Less(i, j int) bool {
	if slice[i].num != slice[j].num {
		return slice[i].num > slice[j].num
	}
	return slice[i].word < slice[j].word
}

func (slice SortSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	var items SortSlice
	for k, v := range m {
		items = append(items, &Item{k, v})
	}
	sort.Sort(items)

	var ans []string
	for i := 0; i < k; i++ {
		ans = append(ans, items[i].word)
	}
	return ans
}

