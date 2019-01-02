func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func getPermutation(n int, k int) string {
	var answer []rune
	used := make([]bool, n + 1)
	k--
	for i := n - 1; i >= 1; i-- {
		f := factorial(i)
		r := k / f
		k = k % f

		t := 0
		for j:=1; j<=n; j++ {
			if used[j] {
				continue
			}
			t++
			if t == r+1 {
				used[j] = true
				answer = append(answer, rune('0' + j))
			}
		}
	}

	for i:=1; i<=n; i++ {
		if used[i] {
			continue
		}
		answer = append(answer, rune('0' + i))
		break
	}

	return string(answer)
}
