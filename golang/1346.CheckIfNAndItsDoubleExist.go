func checkIfExist(arr []int) bool {
    sort.Ints(arr)
    
    binarySearch := func(x int) int {
        l, r := 0, len(arr) - 1
        var mid int
        for l <= r {
            mid = (l+r)/2
            if arr[mid] == x {
                return mid
            }
            if arr[mid] < x {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return -1
    }
    
    var r int
    for i, x := range arr {
        r = binarySearch(x * 2)
        if r != -1 && r != i {
            return true
        }
    }
    return false
}

