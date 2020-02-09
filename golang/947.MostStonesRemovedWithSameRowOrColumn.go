
func removeStones(stones [][]int) int {
    n := len(stones)
    ufs := make([]int, n)
    for i := range ufs {
        ufs[i] = i
    }
    
    var find func(x int) int
    find = func(x int) int {
        if x == ufs[x] {
            return x
        }
        ufs[x] = find(ufs[x])
        return ufs[x]
    }
    
    merge := func(x, y int) {
        x = find(x)
        y = find(y)
        ufs[y] = x
    }

    shareX := make(map[int]int)
    shareY := make(map[int]int)
    for i, stone := range stones {
        x, xe := shareX[stone[0]]
        y, ye := shareY[stone[1]]
        if !xe && !ye {
            shareX[stone[0]] = i
            shareY[stone[1]] = i
            continue
        }
        if xe && ye {
            merge(x, y)
            ufs[i] = x
            continue
        }
        if xe {
            ufs[i] = x
            shareY[stone[1]] = x
        }
        if ye {
            ufs[i] = y
            shareX[stone[0]] = y
        }
    }
    
    m := make(map[int]int)
    for i := range stones {
        m[find(i)]++
    }
    
    ans := 0
    for _, v := range m {
        ans += v-1
    }
    return ans
}

