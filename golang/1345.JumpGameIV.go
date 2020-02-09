func minJumps(arr []int) int {
    reach := make(map[int][]int)
    for i, x := range arr {
        reach[x] = append(reach[x], i)
    }
    for _, v := range reach {
        sort.Sort(sort.Reverse(sort.IntSlice(v)))
    }
    n := len(arr)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = math.MaxInt32
    }
    dp[0] = 0
    
    queue := list.New()
    queue.PushBack(0)
    flag := make([]bool, n)
    
    for queue.Len() > 0 {
        now := queue.Front().Value.(int)
        queue.Remove(queue.Front())
        flag[now] = false

        val := dp[now] + 1
        nxt := now - 1
        if nxt >= 0 {
            update(queue, dp, flag, nxt, val)
        }
        nxt = now + 1
        if nxt < n {
            update(queue, dp, flag, nxt, val)
            if nxt == n-1 {
                return dp[n-1]
            }
        }
        for _, nxt = range reach[arr[now]] {
            update(queue, dp, flag, nxt, val)
            if nxt == n-1 {
                return dp[n-1]
            }
        }
    }
    return -1
}

func update(queue *list.List, dp []int, flag []bool, nxt, val int) {
    if dp[nxt] <= val {
        return
    }
    dp[nxt] = val
    if flag[nxt] {
        return
    }
    queue.PushBack(nxt)
    flag[nxt] = true
}

