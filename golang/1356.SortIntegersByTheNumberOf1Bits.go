type IntSlice []int

func (intSlice IntSlice) Len() int {
    return len(intSlice)
}

func (intSlice IntSlice) Less(i, j int) bool {
    a, b := getOneBit(intSlice[i]), getOneBit(intSlice[j])
    if a != b {
        return a < b
    }
    return intSlice[i] < intSlice[j]
}

func (intSlice IntSlice) Swap(i, j int) {
    intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
}

func getOneBit(x int) int {
    a := 0
    for x > 0 {
        a += x & 1
        x >>= 1
    }
    return a
}

func sortByBits(arr []int) []int {
    sort.Sort(IntSlice(arr))
    return arr
}

