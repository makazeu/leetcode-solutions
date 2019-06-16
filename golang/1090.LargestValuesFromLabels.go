type Item struct {
	value, label int
}

type MyInterface []*Item

func (a MyInterface) Len() int {
	return len(a)
}

func (a MyInterface) Less(i, j int) bool {
	return a[i].value > a[j].value
}

func (a MyInterface) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	var items []*Item
	for i := 0; i < len(values); i++ {
		items = append(items, &Item{values[i], labels[i]})
	}
	sort.Sort(MyInterface(items))

	num := 0
	ans := 0
	m := make(map[int]int)
	for _, item := range items {
		if m[item.label] == use_limit {
			continue
		}
		if num == num_wanted {
			break
		}
		num++
		ans += item.value
		m[item.label] ++
	}
	return ans
}

