func minFlips(a int, b int, c int) int {
	aBin := getBinaryRep(a)
	bBin := getBinaryRep(b)
	cBin := getBinaryRep(c)

	maxLen := len(aBin)
	if len(bBin) > maxLen {
		maxLen = len(bBin)
	}
	if len(cBin) > maxLen {
		maxLen = len(cBin)
	}
	for i := len(aBin) + 1; i <= maxLen; i++ {
		aBin = append(aBin, 0)
	}
	for i := len(bBin) + 1; i <= maxLen; i++ {
		bBin = append(bBin, 0)
	}
	for i := len(cBin) + 1; i <= maxLen; i++ {
		cBin = append(cBin, 0)
	}

	ans := 0
	for i := 0; i < maxLen; i++ {
		if cBin[i] == 0 {
			ans += (aBin[i] + bBin[i])
		} else if cBin[i] == 1 {
			if aBin[i]+bBin[i] == 0 {
				ans++
			}
		}
	}
	return ans
}

func getBinaryRep(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%2)
		num /= 2
	}
	return res
}

