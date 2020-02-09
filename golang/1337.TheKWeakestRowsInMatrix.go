type Item struct {
    num int
    row int
}

type ItemSlice []Item

func (itemSlice ItemSlice) Len() int {
    return len(itemSlice)
}

func (itemSlice ItemSlice) Less(i, j int) bool {
    if itemSlice[i].num != itemSlice[j].num {
        return itemSlice[i].num < itemSlice[j].num
    }
    return itemSlice[i].row < itemSlice[j].row
}

func (itemSlice ItemSlice) Swap(i, j int) {
    itemSlice[i], itemSlice[j] = itemSlice[j], itemSlice[i]
}

func kWeakestRows(mat [][]int, k int) []int {
    var items ItemSlice
    for i, row := range mat {
        num := 0
        for _, x := range row {
            if x == 1 {
                num++
            }
        }
        items = append(items, Item{num, i})
    }
    sort.Sort(items)
    
    var ans []int
    for i := 0; i < k; i++{
        ans = append(ans, items[i].row)
    }
    return ans
}

