func commonChars(A []string) []string {
	var ans []string
	if len(A) == 0 {
		return ans
	}
	num := make([][]int, len(A))
	for i, s := range A {
		num[i] = make([]int, 26)
		for _, c := range s {
			num[i][int(c-'a')]++
		}
	}
	for i := 0; i < 26; i++ {
		n := num[0][i]
		for j := 1; j < len(A); j++ {
			n = min(n, num[j][i])
		}
		for j := 1; j <= n; j++ {
			ans = append(ans, string(i+'a'))
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

