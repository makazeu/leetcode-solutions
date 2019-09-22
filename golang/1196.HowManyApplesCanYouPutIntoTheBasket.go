func maxNumberOfApples(arr []int) int {
    sort.Ints(arr)
    ans := 0
    sum := 0
    for _, w := range arr {
        if(sum + w > 5000) {
            break
        }
        ans ++
        sum += w
    } 
    return ans
}

