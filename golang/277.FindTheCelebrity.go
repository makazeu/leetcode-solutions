/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
    return func(n int) int {
        if n == 0 {
            return -1
        }
        if n == 1 {
            return 0
        }
        flag := make([]bool, n)
        known := make([]int, n)

        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if i == j {
                    continue
                }
                if knows(i, j) {
                    flag[i] = true
                    known[j]++
                } else {
                    flag[j] = true
                }
            }
        }
        fmt.Println(flag)

        for i := 0; i < n; i++ {
            if !flag[i] && known[i] == n-1 {
                return i
            }
        }
        return -1
    }
}

