func longestMountain(A []int) int {
    n := len(A)
    left := make([]int, n)
    right := make([]int, n)
    
    for i:=1; i<n; i++ {
        if A[i] > A[i-1] {
            left[i] = left[i-1] + 1
        }
    }
    for i:=n-2; i>=0; i-- {
        if A[i] > A[i+1] {
            right[i] = right[i+1]+1
        }
    }
    
    ans := 0
    for i:=1; i<=n-2; i++ {
        if right[i] > 0 && left[i] > 0 && right[i] + left[i] + 1 > ans {
            ans = right[i] + left[i] + 1
        }
    }
    return ans
}

