func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	h := make(map[string]int)
	for i, e := range req_skills {
		h[e] = i
	}
	var skills []uint
	for _, e := range people {
		s := uint(0)
		for _, k := range e {
			if a, exists := h[k]; exists {
				s |= 1 << uint(a)
			}
		}
		skills = append(skills, s)
	}

	expected := (1 << uint(len(req_skills))) - 1
	less := len(people)
	var team, ans []int
	best := make(map[uint]int)
	dfs(0, uint(expected), 0, &less, &team, &ans, skills, &best)
	return ans
}

func dfs(skill, expected uint, pos int, less *int, team, ans *[]int, skills []uint, best *map[uint]int) {
	if pos == len(skills) {
		return
	}

	// not grouping current one
	dfs(skill, expected, pos+1, less, team, ans, skills, best)

	// group current one
	newSkill := skill | skills[pos]
	if newSkill == skill {
		return
	}
	if (*best)[newSkill] != 0 && (*best)[newSkill] <= len(*team)+1 {
		return
	}
	(*best)[newSkill] = len(*team) + 1
	*team = append(*team, pos)

	if newSkill == expected {
		if len(*team) < *less {
			*less = len(*team)
			*ans = append([]int(nil), *team...)
		}
	} else {
		dfs(newSkill, expected, pos+1, less, team, ans, skills, best)
	}
	*team = (*team)[:len(*team)-1]
}

