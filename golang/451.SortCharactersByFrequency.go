type Item struct {
	char rune
	num  int
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

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	var items SortSlice
	for k, v := range m {
		items = append(items, &Item{k, v})
	}
	sort.Sort(items)

	var ans []rune
	for _, item := range items {
		for i := 0;i<item.num;i++ {
			ans = append(ans, item.char)
		}
	}
	return string(ans)
}

