func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum % k != 0 {
        return false
    }
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    
    n := len(nums)
    target := sum / k
    used := make([]bool, n)

    found := false
    var dfs func(curSum, curCount, curNum int) 
    dfs = func(curSum, curCount, curNum int) {
        if curCount == k {
            if curNum == n {
                found = true
            }
            return
        }
        for i, num := range nums {
            if used[i] {
                continue
            }
            if curSum + num > target {
                continue
            }
            used[i] = true
            curSum += num
            if curSum == target {
                dfs(0, curCount + 1, curNum + 1)
            } else {
                dfs(curSum, curCount, curNum + 1)
            }
            curSum -= num
            used[i] = false
            
            if found {
                return
            }
        }
    }
    
    dfs(0, 0, 0)
    return found
}

