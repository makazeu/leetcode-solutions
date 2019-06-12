func dfs(ans *int, pos, len int, m map[int]int) {
	if pos > len {
		*ans++
		return
	}

	for k, v := range m {
		if v > 0 {
			m[k]--
			dfs(ans, pos+1, len, m)
			m[k]++
		}
	}
}

func numTilePossibilities(tiles string) int {
	m := make(map[int]int)
	for _, e := range tiles {
		m[int(e-'A')]++
	}
	ans := 0
	for i := 1; i <= len(tiles); i++ {
		dfs(&ans, 1, i, m)
	}
	return ans
}

