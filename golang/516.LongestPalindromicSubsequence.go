func longestPalindromeSubseq(s string) int {
    l := len(s)
    dp := make([][]int, l)
    for i := range dp {
        dp[i] = make([]int, l)
    }
    
    for k := 1; k<=l ; k++ {
        for i:=0; i<l; i++ {
            j := i+k-1
            if j >= l {
                break
            }
            if i == j {
                dp[i][j] = 1
                continue
            }
            if dp[i+1][j] > dp[i][j] {
                dp[i][j] = dp[i+1][j]
            }
            if dp[i][j-1] > dp[i][j] {
                dp[i][j] = dp[i][j-1]
            }
            if s[i] == s[j] && dp[i+1][j-1] + 2 > dp[i][j] {
                dp[i][j] = dp[i+1][j-1] + 2
            }
        }
    }
    
    return dp[0][l-1]
}

