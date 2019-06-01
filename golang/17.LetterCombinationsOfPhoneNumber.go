var NumberMapping = [][]rune{
	{}, {},               // 0 , 1
	{'a', 'b', 'c'},      // 2
	{'d', 'e', 'f'},      // 3
	{'g', 'h', 'i'},      // 4
	{'j', 'k', 'l'},      // 5
	{'m', 'n', 'o'},      // 6
	{'p', 'q', 'r', 's'}, // 7
	{'t', 'u', 'v'},      // 8
	{'w', 'x', 'y', 'z'}, // 9
}

func dfs(digits *string, pos int, current *[]rune, result *[]string) {
	if pos == len(*digits) {
		*result = append(*result, string(*current))
		return
	}
	for _, e := range NumberMapping[(*digits)[pos]-'0'] {
		(*current)[pos] = e
		dfs(digits, pos + 1, current, result)
	}
}

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}
	current := make([]rune, len(digits))
	dfs(&digits, 0, &current, &result)
	return result
}

