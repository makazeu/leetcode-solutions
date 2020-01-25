func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    m := make(map[int]int)
    queue := list.New()
    inQueue := make(map[int]bool)
    for _, coin := range coins {
        m[coin] = 1
        queue.PushBack(coin)
        inQueue[coin] = true
    }
    if m[amount] == 1 {
        return 1
    }
    
    for queue.Len() > 0 {
        cnt := queue.Front().Value.(int)
        queue.Remove(queue.Front())
        inQueue[cnt] = false
        cost := m[cnt]
        
        var nxt int
        for _, coin := range coins {
            nxt = cnt + coin
            if nxt == amount {
                return cost + 1
            }
            if nxt > amount {
                continue
            }
            if preCost, exists := m[nxt]; !exists || preCost > cost + 1 {
                m[nxt] = cost + 1
                if !inQueue[nxt] {
                    queue.PushBack(nxt)
                    inQueue[nxt] = true
                }
            }
        }
    }
    return -1
}

