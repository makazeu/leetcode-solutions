func minSetSize(arr []int) int {
    hash := make(map[int]int)
    for _, x := range arr {
        hash[x]++
    }
    
    var nums []int
    for _, v := range hash {
        nums = append(nums, v)
    }
    
    sort.Ints(nums)
    l := len(nums)
    now := 0
    for i := l-1; i >= 0; i-- {
        now += nums[i]
        if now >= len(arr)/2 {
            return l-i
        }
    }
    return -1
}

