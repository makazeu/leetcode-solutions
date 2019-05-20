type Interface struct {
	words []string
}

func (a Interface) Len() int {
	return len(a.words)
}

func (a Interface) Less(i, j int) bool {
	return len(a.words[i]) < len(a.words[j])
}

func (a Interface) Swap(i, j int) {
	a.words[i], a.words[j] = a.words[j], a.words[i]
}

func check(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}

	i, j := 0, 0
	flag := false

	for i < len(a) && j < len(b) {
		if a[i] != b[j] {
			if flag {
				return false
			}
			j++
			flag = true
		} else {
			i++
			j++
		}
	}
	return true
}

func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}
	if len(words) == 1 {
		return 1
	}
	var w Interface
	for _, e := range words {
		w.words = append(w.words, e)
	}
	sort.Sort(w)

	ans := 0
	dp := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		dp[i] = 0
	}
	for i := 0; i < len(words); i++ {
		a := w.words[i]
		for j := 0; j < i; j++ {
			b := w.words[j]
			if len(a)-1 != len(b) {
				continue
			}

			if check(b, a) && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
			if dp[i] > ans {
				ans = dp[i]
			}
		}
	}
	return ans + 1
}

