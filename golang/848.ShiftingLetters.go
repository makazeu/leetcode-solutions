const MOD = 26

func shiftingLetters(S string, shifts []int) string {
    for i := len(shifts) - 2; i>=0; i-- {
        shifts[i] = (shifts[i] + shifts[i+1]) % MOD
    }
    var result []rune
    for i := range S {
        result = append(result, rune('a' + (int(S[i]-'a') + shifts[i]) % MOD))
    }
    return string(result)
}

