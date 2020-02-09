func numberOfSteps (num int) int {
    step := 0
    for ; num > 0; step++ {
        if num & 1 == 1 {
            num--
        } else {
            num >>= 1
        }
    }
    return step
}

