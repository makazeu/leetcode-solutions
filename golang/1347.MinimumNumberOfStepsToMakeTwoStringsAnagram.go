func minSteps(s string, t string) int {
    sh := make(map[int]int)
    th := make(map[int]int)
    
    for _, c := range s {
        sh[int(c-'a')]++
    }
    for _, c := range t {
        th[int(c-'a')]++
    }
    
    ans := 0
    var tmp int
    for k, v := range th {
        tmp = v - sh[k]
        if tmp > 0 {
            ans += tmp
        }
    } 
    return ans
}

