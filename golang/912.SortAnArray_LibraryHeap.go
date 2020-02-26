type IntSlice []int

func (intSlice IntSlice) Len() int {
    return len(intSlice)
}

func (intSlice IntSlice) Less(i,j int) bool {
    return intSlice[i] < intSlice[j]
}

func (intSlice IntSlice) Swap(i, j int) {
    intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
}

func (intSlice *IntSlice) Push(x interface{}) {
    *intSlice = append(*intSlice, x.(int))
}

func (intSlice *IntSlice) Pop() interface{} {
    l := intSlice.Len() - 1
    e := (*intSlice)[l]
    (*intSlice) = (*intSlice)[:l]
    return e
}

func sortArray(nums []int) []int {
    h := IntSlice(nums)
    heap.Init(&h)

    n := len(nums)
    var ans []int
    for i := 0; i < n; i++ {
        ans = append(ans, heap.Pop(&h).(int))
    }
    return ans
}

