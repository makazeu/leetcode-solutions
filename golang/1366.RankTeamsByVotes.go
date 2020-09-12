type Item struct {
	count []int
	char  rune
}

type ItemSlice []Item

func (itemSlice ItemSlice) Len() int {
	return len(itemSlice)
}

func (itemSlice ItemSlice) Less(i, j int) bool {
	for a := range itemSlice[i].count {
		if itemSlice[i].count[a] != itemSlice[j].count[a] {
			return itemSlice[i].count[a] > itemSlice[j].count[a]
		}
	}
	return itemSlice[i].char < itemSlice[j].char
}

func (itemSlice ItemSlice) Swap(i, j int) {
	itemSlice[i], itemSlice[j] = itemSlice[j], itemSlice[i]
}

func rankTeams(votes []string) string {
	n := len(votes[0])
	pos := make(map[rune]Item)
	for _, c := range votes[0] {
		pos[c] = Item{make([]int, n), c}
	}
	for _, vote := range votes {
		for i, c := range vote {
			pos[c].count[i]++
		}
	}

	var itemSlice ItemSlice
	for _, item := range pos {
		itemSlice = append(itemSlice, item)
	}
	sort.Sort(itemSlice)

	var ans string
	for _, item := range itemSlice {
		ans += string(item.char)
	}
	return ans
}

