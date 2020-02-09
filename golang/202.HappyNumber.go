func isHappy(n int) bool {
    if n == 1 {
        return true
    }
    hash := make(map[int]bool)
    hash[n] = true
    
    var now, m int
    for {
        now = 0
        for n > 0 {
            m = n % 10
            now += m * m
            n /= 10
        }
        if now == 1 {
            return true
        }
        if hash[now] {
            break
        }
        hash[now] = true
        n = now
    }
    return false
}

