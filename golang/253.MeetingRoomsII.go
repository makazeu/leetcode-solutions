const End = 0
const Start = 1

type Item struct {
    pos, event int
}

type ItemSlice []Item

func (itemSlice ItemSlice) Len() int {
    return len(itemSlice)
}

func (itemSlice ItemSlice) Less(i,j int) bool {
    if itemSlice[i].pos != itemSlice[j].pos {
        return itemSlice[i].pos < itemSlice[j].pos
    }
    return itemSlice[i].event < itemSlice[j].event
}

func (itemSlice ItemSlice) Swap(i, j int) {
    itemSlice[i], itemSlice[j] = itemSlice[j], itemSlice[i]
}

func minMeetingRooms(intervals [][]int) int {
    var itemSlice ItemSlice
    for _, itvl := range intervals {
        itemSlice = append(itemSlice, Item{itvl[0], Start})
        itemSlice = append(itemSlice, Item{itvl[1], End})
    }
    sort.Sort(itemSlice)

    now := 0
    max := 0
    for _, item := range itemSlice {
        if item.event == End {
            now--
        } else {
            now++
            if now > max {
                max = now
            }
        }
    }
    return max
}

