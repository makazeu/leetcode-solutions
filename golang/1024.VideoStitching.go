func videoStitching(clips [][]int, T int) int {
	ans := make([]int, T+1)
	for i := range ans {
		ans[i] = -1
	}
	ans[0] = 0

	for t := 1; t <= T; t++ {
		for i := range clips {
			if clips[i][0] >= t || clips[i][1] < t {
				continue
			}

			for j := clips[i][0]; j < t; j++ {
				if ans[j] != -1 && (ans[t] == -1 || ans[j]+1 < ans[t]) {
					ans[t] = ans[j] + 1
				}
			}
		}
	}

	return ans[T]
}

