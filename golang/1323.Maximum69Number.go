func maximum69Number (num int) int {
    var digits []int
    for num > 0 {
        digits = append(digits, num % 10)
        num /= 10
    }
    
    for i:=len(digits)-1; i>=0; i-- {
        if digits[i] == 6 {
            digits[i] = 9
            break
        }
    }
    
    num = 0
    for i:=len(digits)-1; i>=0; i-- {
        num = num * 10 + digits[i]
    }
    return num
}

