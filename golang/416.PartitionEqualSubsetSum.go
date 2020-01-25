func canPartition(nums []int) bool {
    sum := 0
    l := len(nums)
    for _, num := range nums {
        sum += num
    }
    if sum % 2 != 0 {
        return false
    }
    sum /= 2
    
    dp := make([][]bool, l + 1)
    for i := range dp {
        dp[i] = make([]bool, sum + 1)
    }
    dp[0][0] = true
    
    for i:=1; i<=l; i++ {
        for j := 1; j <= sum; j++ {
            if dp[i-1][j] {
                dp[i][j] = true
                continue
            }
            if nums[i-1] > j {
                continue
            }
            if dp[i-1][j-nums[i-1]] {
                dp[i][j] = true
            }
        }
    }
    
    return dp[l][sum]
}

