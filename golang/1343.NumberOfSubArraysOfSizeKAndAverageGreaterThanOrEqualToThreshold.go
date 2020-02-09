func numOfSubarrays(arr []int, k int, threshold int) int {
    threshold *= k
    sum := 0
    ans := 0
    for i:=0; i<k; i++ {
        sum += arr[i]
    }
    if sum >= threshold {
        ans++
    }
    
    for i:=k; i<len(arr); i++ {
        sum += arr[i]
        sum -= arr[i-k]
        if sum >= threshold {
            ans++
        }
    }
    return ans
}

