type Item struct {
    val, pos int
}

type ItemSlice []Item

func (itemSlice ItemSlice) Len() int {
    return len(itemSlice)
}

func (itemSlice ItemSlice) Less(i, j int) bool {
    if itemSlice[i].val != itemSlice[j].val {
        return itemSlice[i].val < itemSlice[j].val
    }
    return itemSlice[i].pos < itemSlice[j].pos
}

func (itemSlice ItemSlice) Swap(i, j int) {
    itemSlice[i], itemSlice[j] = itemSlice[j], itemSlice[i]
}

func maxJumps(arr []int, d int) int {
    var itemSlice ItemSlice
    for i, x := range arr {
        itemSlice = append(itemSlice, Item{x, i})
    }
    sort.Sort(itemSlice)
    
    dp := make([]int, len(arr))
    ans := 0
    for _, item := range itemSlice {
        pos := item.pos
        for i := 1; i <= d && pos-i >= 0; i++ {
            if arr[pos] <= arr[pos-i] {
                break
            }
            if dp[pos] < dp[pos-i] {
                dp[pos] = dp[pos-i]
            }
        }
        
        for i := 1; i <= d && pos+i < len(arr); i++ {
            if arr[pos] <= arr[pos+i] {
                break
            }
            if dp[pos] < dp[pos+i] {
                dp[pos] = dp[pos+i]
            }
        }
        dp[pos]++
        if dp[pos] > ans {
            ans = dp[pos]
        }
    }
    return ans
}

