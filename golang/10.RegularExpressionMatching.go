func isMatch(s string, p string) bool {
	var splitPattern []string
	for i := 0; i < len(p); i++ {
		if i >= len(p)-1 || p[i+1] != '*' {
			splitPattern = append(splitPattern, p[i:i+1])
		} else {
			splitPattern = append(splitPattern, p[i:i+2])
			i++
		}
	}

	res := make([][]bool, len(splitPattern)+1)
	for i := range res {
		res[i] = make([]bool, len(s)+1)
	}
	res[0][0] = true
	for pos := 1; pos <= len(splitPattern); pos++ {
		for i := 1; i <= len(s); i++ {
			if res[pos][i] {
				continue
			}
			if !res[pos-1][i-1] {
				continue
			}
			if len(splitPattern[pos-1]) == 1 {
				if s[i-1] == splitPattern[pos-1][0] || splitPattern[pos-1] == "." {
					res[pos][i] = true
				}
				continue
			}
			res[pos][i-1] = res[pos-1][i-1]
			for j := i; j <= len(s); j++ {
				if s[j-1] == splitPattern[pos-1][0] || splitPattern[pos-1][0] == '.' {
					res[pos][j] = true
				} else {
					break
				}
			}
		}
		//fmt.Println(res[pos])
	}

	for i := len(splitPattern); i >= 0; i-- {
		if res[i][len(s)] {
			return true
		}
		if i >= 1 && len(splitPattern[i-1]) == 1 {
			break
		}
	}

	return false
}

