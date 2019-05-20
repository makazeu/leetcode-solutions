func removeDuplicates(S string) string {
	if len(S) == 0 {
		return S
	}

	pre := make([]int, len(S))
	deleted := make([]bool, len(S))

	for i := 0; i < len(S); i++ {
		pre[i] = i - 1
	}
	pre[0] = -1

	for i := 0; i < len(S); i++ {
		if pre[i] == -1 {
			continue
		}
		if S[i] != S[pre[i]] {
			continue
		}
		deleted[i] = true
		deleted[pre[i]] = true
		if i < len(S)-1 {
			p := pre[pre[i]]
			pre[i+1] = p
		}
	}

	var ans []byte
	for i := 0; i < len(S); i++ {
		if !deleted[i] {
			ans = append(ans, S[i])
		}
	}
	return string(ans)
}

