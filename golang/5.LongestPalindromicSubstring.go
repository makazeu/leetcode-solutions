func longestPalindrome(s string) string {
    maxLen := 0
    var maxStr string
    l := len(s)
    dp := make([][]bool, l)
    for i := range dp {
        dp[i] = make([]bool, l)
    }

    for k := 1; k <= l; k++ {
        for i := 0; i <= l-k; i++ {
            j := i+k-1
            if s[i] != s[j] {
                continue
            }
            if j-i <= 1 {
                dp[i][j] = true
            } else if dp[i+1][j-1] {
                dp[i][j] = true
            }

            if dp[i][j] && j-i+1 > maxLen {
                maxLen = j-i+1
                maxStr = s[i:j+1]
            }
        }
    }
    return maxStr
}

