func longestCommonSubsequence(text1 string, text2 string) int {
    l1 := len(text1)
    l2 := len(text2)
    dp := make([][]int, l1 + 1)
    dp[0] = make([]int, l2 + 1)
    for i:=1; i<=l1; i++ {
        dp[i] = make([]int, l2 + 1)
        for j:=1; j<=l2; j++ {
            if dp[i-1][j] > dp[i][j] {
                dp[i][j] = dp[i-1][j]
            }
            if dp[i][j-1] > dp[i][j] {
                dp[i][j] = dp[i][j-1]
            }
            if dp[i-1][j-1] + 1 > dp[i][j] && text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            }
        }
    }
    return dp[l1][l2]
}

