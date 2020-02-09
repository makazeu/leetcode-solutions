func minSubArrayLen(s int, nums []int) int {
    slow, fast := -1, -1
    ans := math.MaxInt32
    n := len(nums)
    sum := 0
    for fast < n-1{
        fast++
        sum += nums[fast]
        
        if sum >= s {
            for sum >= s {
                slow++
                sum -= nums[slow]
            }
            if ans > fast-slow+1 {
                ans = fast-slow + 1
            }
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return ans
}
