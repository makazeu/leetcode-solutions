type Account struct {
	name   string
	email  string
	parent int
}

func find(ufs []*Account, i int) int {
	if ufs[i].parent != i {
		ufs[i].parent = find(ufs, ufs[i].parent)
	}
	return ufs[i].parent
}

func union(ufs []*Account, i, p int) {
	i = find(ufs, i)
	p = find(ufs, p)
	if i != p {
		ufs[i].parent = p
	}
}

func accountsMerge(accounts [][]string) [][]string {
	m := make(map[string]int)
	var ufs []*Account
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			if _, exists := m[email]; !exists {
				l := len(ufs)
				ufs = append(ufs, &Account{name, email, l})
				m[email] = l
			}
		}
	}

	for _, account := range accounts {
		if len(account) <= 1 {
			continue
		}
		p := find(ufs, m[account[1]])
		for i := 2; i < len(account); i++ {
			union(ufs, m[account[i]], p)
		}
	}

	ansMap := make(map[int][]string)
	for i, account := range ufs {
		p := find(ufs, i)
		ansMap[p] = append(ansMap[p], account.email)
	}

	var result [][]string
	for k, v := range ansMap {
		sort.Strings(v)
		result = append(result, []string{ufs[k].name})
		l := len(result) - 1
		result[l] = append(result[l], v...)
	}

	return result
}

